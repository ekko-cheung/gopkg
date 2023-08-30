package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/veerdone/gopkg/log"
	"go.uber.org/zap"
)

type RClient struct {
	*redis.Client
	log log.Log
}

func NewRClient(client *redis.Client, log log.Log) *RClient {
	return &RClient{Client: client, log: log}
}

func (r *RClient) Get(ctx context.Context, key string) *redis.StringCmd {
	cmd := r.Client.Get(ctx, key)
	err := cmd.Err()
	if err != nil && err != redis.Nil {
		r.log.Warn(ctx, "get redis cache fail", zap.Error(err), zap.String("key", key))
	}

	return cmd
}

func (r *RClient) Scan(ctx context.Context, key string, dest interface{}) error {
	err := r.Client.Get(ctx, key).Scan(dest)
	if err != nil {
		if err != redis.Nil {
			r.log.Warn(ctx, "get redis cache fail or scan value to redis fail", zap.Error(err), zap.String("key", key))
		}
		return err
	}

	return nil
}
