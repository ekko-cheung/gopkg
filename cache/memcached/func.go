package memcached

import (
	"encoding/binary"
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/veerdone/gopkg/log"
	"github.com/veerdone/gopkg/util"
	"go.uber.org/zap"
)

var (
	marshal   = json.Marshal
	unMarshal = json.Unmarshal
)

func SetMarshalAndUnMarshal(m func(interface{}) ([]byte, error), un func([]byte, interface{}) error) {
	marshal = m
	unMarshal = un
}

func Set(item *memcache.Item) {
	if err := mc.Set(item); err != nil {
		log.Warn("memcached set cache error", zap.Error(err), zap.String("key", item.Key))
	}
}

func SetStr(key, val string) {
	SetStrEx(key, val, 0)
}

func SetStrEx(key, val string, expire int32) {
	Set(&memcache.Item{
		Key:        key,
		Value:      util.StringToSliceByte(val),
		Expiration: expire,
	})
}

func SetByte(key string, val []byte) {
	SetByteEx(key, val, 0)
}

func SetByteEx(key string, val []byte, expire int32) {
	Set(&memcache.Item{
		Key:        key,
		Value:      val,
		Expiration: expire,
	})
}

func SetStruct(key string, val interface{}) {
	SetStructEx(key, val, 0)
}

func SetStructEx(key string, val interface{}, expire int32) {
	bytes, err := marshal(val)
	if err != nil {
		log.Warn("marshal struct error", zap.Error(err), zap.String("key", key))
		return
	}
	Set(&memcache.Item{
		Key:        key,
		Value:      bytes,
		Expiration: expire,
	})
}

func SetInt(key string, i int64) {
	SetIntEx(key, i, 0)
}

func SetIntEx(key string, i int64, expire int32) {
	b := make([]byte, 8)
	binary.PutVarint(b, i)
	Set(&memcache.Item{
		Key:        key,
		Value:      b,
		Expiration: expire,
	})
}

func SetUint(key string, i uint64) {
	SetUintEx(key, i, 0)
}

func SetUintEx(key string, i uint64, expire int32) {
	b := make([]byte, 8)
	binary.PutUvarint(b, i)
	Set(&memcache.Item{
		Key:        key,
		Value:      b,
		Expiration: expire,
	})
}

func Increment(key string, i uint64) uint64 {
	value, err := mc.Increment(key, i)
	if err != nil {
		log.Warn("memcached interment error", zap.Error(err))
	}

	return value
}

func Decrement(key string, i uint64) uint64 {
	value, err := mc.Decrement(key, i)
	if err != nil {
		log.Warn("memcached decrement error", zap.Error(err))
	}

	return value
}

func GByte(key string) ([]byte, error) {
	item, err := mc.Get(key)
	if err != nil {
		return nil, err
	}
	return item.Value, nil
}

func GStr(key string) (string, error) {
	bytes, err := GByte(key)
	if err != nil {
		if err != memcache.ErrCacheMiss {
			log.Warn("memcached get cache error", zap.Error(err))
		}
		return "", err
	}

	return util.SliceByteToString(bytes), nil
}

func GInt(key string) (int64, error) {
	bytes, err := GByte(key)
	if err != nil {
		if err != memcache.ErrCacheMiss {
			log.Warn("memcached get cache error", zap.Error(err))
		}
		return 0, err
	}
	i, _ := binary.Varint(bytes)
	return i, nil
}

func GUint(key string) (uint64, error) {
	bytes, err := GByte(key)
	if err != nil {
		if err != memcache.ErrCacheMiss {
			log.Warn("memcached get cache error", zap.Error(err))
		}
		return 0, err
	}
	i, _ := binary.Uvarint(bytes)
	return i, nil
}

func GStruct(key string, v interface{}) error {
	bytes, err := GByte(key)
	if err != nil {
		log.Warn("memcached get cache error", zap.Error(err))
		return err
	}

	return unMarshal(bytes, v)
}

func Del(key string) {
	if err := mc.Delete(key); err != nil {
		log.Warn("memcached del cache error", zap.Error(err))
	}
}

func DelAll() {
	if err := mc.DeleteAll(); err != nil {
		log.Warn("memcached del all cache error", zap.Error(err))
	}
}

func CAS(item *memcache.Item) {
	if err := mc.CompareAndSwap(item); err != nil {
		log.Warn("memcached CAS set cache error", zap.Error(err), zap.String("key", item.Key))
	}
}

func CASStruct(key string, v interface{}) {
	CASStructEx(key, v, 0)
}

func CASStructEx(key string, v interface{}, expire int32) {
	bytes, err := marshal(v)
	if err != nil {
		log.Warn("marshal struct error", zap.Error(err))
		return
	}
	CAS(&memcache.Item{
		Key:        key,
		Value:      bytes,
		Expiration: expire,
	})
}

func CASInt(key string, i int64) {
	CASIntEx(key, i, 0)
}

func CASIntEx(key string, i int64, expire int32) {
	b := make([]byte, 8)
	binary.PutVarint(b, i)
	CAS(&memcache.Item{
		Key:        key,
		Value:      b,
		Expiration: expire,
	})
}

func CASUint(key string, i uint64) {
	CASUintEx(key, i, 0)
}

func CASUintEx(key string, i uint64, expire int32) {
	b := make([]byte, 8)
	binary.PutUvarint(b, i)
	CAS(&memcache.Item{
		Key:        key,
		Value:      b,
		Expiration: expire,
	})
}
