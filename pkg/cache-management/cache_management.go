package cacheManagement

import (
	"github.com/go-redis/redis/v8"
	"strings"
	"sync"
	"url-shortener/pkg/config"
)

type CacheManagement interface {
	SetString(key string, value string, expire int) error
	Set(key string, value interface{}, expire int) error
	Sets(values map[string]string, expire int) error
	Get(key string) (*string, error)
	Delete(key string) error
}

var (
	once            sync.Once
	dictionaryCache CacheManagement
	redisCache      CacheManagement
)

// NewCacheManagerFactory returns a factory function for creating CacheManagement instances.
func NewCacheManagerFactory() func() CacheManagement {
	return func() CacheManagement {
		config := configReader.GetInstance()
		cacheType := config.Cache.Type

		once.Do(func() {
			if strings.ToUpper(cacheType) == strings.ToUpper("redis") {
				redisCache = newRedisCacheManagement()
				dictionaryCache = nil
			} else {
				dictionaryCache = newDictionaryCacheManagement()
				redisCache = nil
			}
		})

		if strings.ToUpper(cacheType) == strings.ToUpper("redis") {
			return redisCache
		} else {
			return dictionaryCache
		}
	}
}

// NewDictionaryCacheManagement creates a new instance of DictionaryCacheManager.
func newDictionaryCacheManagement() *DictionaryCacheManager {
	return &DictionaryCacheManager{
		cache: make(map[string]string),
	}
}

// NewRedisCacheManagement creates a new instance of RedisCacheManager.
func newRedisCacheManagement() *RedisCacheManager {
	config := configReader.GetInstance()
	return &RedisCacheManager{
		client: redis.NewClient(&redis.Options{
			Addr:     config.Cache.Address,
			Username: config.Cache.Username,
			Password: config.Cache.Password,
			DB:       0,
		}),
	}
}
