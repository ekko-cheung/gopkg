package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
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
	cli *redis.Client
	key string
	ctx context.Context
}

func NewLock(key string, ctx context.Context) *Lock {
	return &Lock{
		cli: client,
		key: key,
		ctx: ctx,
	}
}

func (l *Lock) Lock(val interface{}, exTime time.Duration) {
	l.lock(val, exTime)
}

func (l *Lock) TryLock(val interface{}, exTime time.Duration) bool {
	ok, _ := l.lock(val, exTime)

	return ok
}

func (l *Lock) lock(val interface{}, exTime time.Duration) (bool, error) {
	b, err := l.cli.SetNX(l.ctx, l.key, val, exTime).Result()
	if err != nil {
		log.CtxWarn(l.ctx, "redis setNX fail", zap.Error(err))
	}

	return b, err
}

func (l *Lock) UnLock(val interface{}) {
	err := unlockScript.Run(l.ctx, l.cli, []string{l.key}, val).Err()
	if err != nil {
		log.CtxWarn(l.ctx, "redis run unlock lua script fail", zap.Error(err))
	}
}
