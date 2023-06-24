package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/veerdone/gopkg/log"
	"go.uber.org/zap"
	"time"
)

var unlockScript = redis.NewScript(`
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	end
	return 0
`)

type Lock struct {
	cli        *redis.Client
	key        string
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewLock(key string, ctx context.Context) *Lock {
	return &Lock{
		cli: client,
		key: key,
		ctx: ctx,
	}
}

func (l *Lock) Lock(val interface{}, exTime time.Duration) bool {
	ok, err := l.lock(val, exTime)
	if err != nil || !ok {
		return false
	}

	go l.watchDog(exTime)
	return true
}

// TryLock try lock in waitTime, eg:
//
//  l := NewLock(key, ctx)
//  if l.TryLock(val, time.Second, time.Second) {
//    defer l.Unlock(val)
//    // business
//  }
func (l *Lock) TryLock(val interface{}, waitTime, exTime time.Duration) bool {
	ctx, cancel := context.WithTimeout(l.ctx, waitTime)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return false
		default:
			if l.Lock(val, exTime) {
				return true
			}
		}
	}
}

func (l *Lock) lock(val interface{}, exTime time.Duration) (bool, error) {
	b, err := l.cli.SetNX(l.ctx, l.key, val, exTime).Result()
	if err != nil {
		log.CtxWarn(l.ctx, "redis setNX fail", zap.Error(err))
	}

	return b, err
}

func (l *Lock) watchDog(exTime time.Duration) {
	l.ctx, l.cancelFunc = context.WithCancel(l.ctx)
	for {
		select {
		case <-l.ctx.Done():
			return
		default:
			l.cli.Expire(l.ctx, l.key, exTime)
			time.Sleep(exTime / 3)
		}
	}
}

func (l *Lock) UnLock(val interface{}) {
	err := unlockScript.Run(l.ctx, l.cli, []string{l.key}, val).Err()
	if err != nil {
		log.CtxWarn(l.ctx, "redis run unlock lua script fail", zap.Error(err))
	}
	if l.cancelFunc != nil {
		l.cancelFunc()
		l.cancelFunc = nil
	}
}
