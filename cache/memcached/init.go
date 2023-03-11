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
