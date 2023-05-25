package redis

import (
	"context"
	"encoding"
	"github.com/go-redis/redis/v8"
	"github.com/veerdone/gopkg/log"
	"github.com/veerdone/gopkg/util"
	"go.uber.org/zap"
	"time"
)

var Nil = redis.Nil

func Client() *redis.Client {
	return client
}

func Set(ctx context.Context, key string, val interface{}) {
	SetEx(ctx, key, val, 0)
}

func SetEx(ctx context.Context, key string, val interface{}, ex time.Duration) {
	err := client.Set(ctx, key, val, ex).Err()
	if err != nil {
		log.Warn("redis set cache error", zap.Error(err))
	}
}

func get(ctx context.Context, key string) (*redis.StringCmd, error) {
	cmd := client.Get(ctx, key)
	if err := cmd.Err(); err != nil {
		if err != redis.Nil {
			log.Warn("redis get cache error", zap.Error(err))
		}
		return nil, err
	}

	return cmd, nil
}

func GBytes(ctx context.Context, key string) ([]byte, error) {
	cmd, err := get(ctx, key)
	if err != nil {
		return nil, err
	}

	return cmd.Bytes()
}

func GStr(ctx context.Context, key string) (string, error) {
	cmd, err := get(ctx, key)
	if err != nil {
		return "", err
	}
	return cmd.String(), nil
}

func GInt(ctx context.Context, key string) (int, error) {
	cmd, err := get(ctx, key)
	if err != nil {
		return 0, err
	}

	return cmd.Int()
}

func GInt64(ctx context.Context, key string) (int64, error) {
	cmd, err := get(ctx, key)
	if err != nil {
		return 0, err
	}
	return cmd.Int64()
}

func GUint64(ctx context.Context, key string) (uint64, error) {
	cmd, err := get(ctx, key)
	if err != nil {
		return 0, err
	}
	return cmd.Uint64()
}

func GScan(ctx context.Context, key string, val interface{}) error {
	cmd, err := get(ctx, key)
	if err != nil {
		return err
	}

	return cmd.Scan(val)
}

func GStruct(ctx context.Context, key string, val encoding.BinaryUnmarshaler) error {
	cmd, err := get(ctx, key)
	if err != nil {
		return err
	}
	s := cmd.Val()

	return val.UnmarshalBinary(util.StringToSliceByte(s))
}

func Del(ctx context.Context, keys ...string) {
	if err := client.Del(ctx, keys...).Err(); err != nil {
		log.Warn("redis del cache error", zap.Error(err))
	}
}

func Expire(ctx context.Context, key string, ex time.Duration) {
	if err := client.Expire(ctx, key, ex).Err(); err != nil {
		log.Warn("redis set key expire error", zap.Error(err))
	}
}

func ExpireAt(ctx context.Context, key string, ex time.Time) {
	if err := client.ExpireAt(ctx, key, ex).Err(); err != nil {
		log.Warn("redis set key expire at error", zap.Error(err))
	}
}

func Incr(ctx context.Context, key string) {
	IncrBy(ctx, key, 1)
}

func IncrBy(ctx context.Context, key string, i int64) {
	if err := client.IncrBy(ctx, key, i).Err(); err != nil {
		log.Warn("redis increment error", zap.Error(err))
	}
}

func Decr(ctx context.Context, key string) {
	DecrBy(ctx, key, 1)
}

func DecrBy(ctx context.Context, key string, i int64) {
	if err := client.DecrBy(ctx, key, i).Err(); err != nil {
		log.Warn("redis decrement error", zap.Error(err))
	}
}

func TTL(ctx context.Context, key string) (time.Duration, error) {
	ttl := client.TTL(ctx, key)
	if err := ttl.Err(); err != nil {
		if err != redis.Nil {
			log.Warn("redis get cache TTL error", zap.Error(err))
		}
		return 0, err
	}

	return ttl.Val(), nil
}

func SetNX(ctx context.Context, key string, val interface{}, ex time.Duration) {
	if err := client.SetNX(ctx, key, val, ex).Err(); err != nil {
		log.Warn("redis setNX error", zap.Error(err))
	}
}
