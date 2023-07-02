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
