/*
 * Copyright 2023 veerdone
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
	cancelFunc context.CancelFunc
}

func NewLock(key string) *Lock {
	return &Lock{
		cli: client,
		key: key,
	}
}

func (l *Lock) Lock(ctx context.Context, val interface{}, exTime time.Duration) bool {
	ok, err := l.lock(ctx, val, exTime)
	if err != nil || !ok {
		return false
	}

	go l.watchDog(ctx, exTime)
	return true
}

// TryLock try lock in waitTime, eg:
//
//	l := NewLock(key, ctx)
//	if l.TryLock(val, time.Second, time.Second) {
//	  defer l.Unlock(val)
//	  // business
//	}
func (l *Lock) TryLock(ctx context.Context, val interface{}, waitTime, exTime time.Duration) bool {
	ctx, cancel := context.WithTimeout(ctx, waitTime)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return false
		default:
			if l.Lock(ctx, val, exTime) {
				return true
			}
		}
	}
}

func (l *Lock) lock(ctx context.Context, val interface{}, exTime time.Duration) (bool, error) {
	b, err := l.cli.SetNX(ctx, l.key, val, exTime).Result()
	if err != nil {
		log.CtxWarn(ctx, "redis setNX fail", zap.Error(err))
	}

	return b, err
}

func (l *Lock) watchDog(ctx context.Context, exTime time.Duration) {
	c, cancelFunc := context.WithCancel(ctx)
	l.cancelFunc = cancelFunc
	for {
		select {
		case <-c.Done():
			return
		default:
			l.cli.Expire(c, l.key, exTime)
			time.Sleep(exTime / 3)
		}
	}
}

func (l *Lock) UnLock(ctx context.Context, val interface{}) {
	err := unlockScript.Run(ctx, l.cli, []string{l.key}, val).Err()
	if err != nil {
		log.CtxWarn(ctx, "redis run unlock lua script fail", zap.Error(err))
	}
	if l.cancelFunc != nil {
		l.cancelFunc()
		l.cancelFunc = nil
	}
}
