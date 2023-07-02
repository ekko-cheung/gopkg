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

package memcached

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/veerdone/gopkg/conf"
	"github.com/veerdone/gopkg/log"
	"go.uber.org/zap"
)

var mc *memcache.Client

func InitMemcached(conf *conf.Memcached) {
	address := conf.Address
	if len(address) == 0 {
		log.Warn("none config for memcached")
		return
	}
	mc = memcache.New(address...)
	err := mc.Ping()
	if err != nil {
		log.Error("ping memcached error", zap.Error(err))
	}
}
