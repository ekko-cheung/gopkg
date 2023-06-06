package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/veerdone/gopkg/conf"
	"github.com/veerdone/gopkg/log"
	"go.uber.org/zap"
)

var client *redis.Client

func InitRedis(conf conf.Redis) {
	if conf.Addr == "" {
		log.Warn("none config for redis")
		return
	}

	client = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Pass,
		Username: conf.Username,
		DB:       conf.Db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Error("ping redis error", zap.Error(err))
	}
}
