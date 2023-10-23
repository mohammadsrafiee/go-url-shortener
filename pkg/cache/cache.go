package cache

import (
	"github.com/go-redis/redis/v8"
	"strings"
	"sync"
	"url-shortener/pkg/config-reader"
)

type Management interface {
	SetString(key string, value string, expire int) error
	Set(key string, value interface{}, expire int) error
	Sets(values map[string]string, expire int) error
	Get(key string) (*string, error)
	Delete(key string) error
}

var (
	once            sync.Once
	dictionaryCache Management
	redisCache      Management
)

func Instance() Management {
	config := configReader.Instance()
	cacheType := config.Cache.Type
	if strings.ToUpper(cacheType) == strings.ToUpper("redis") {
		return redisCache
	} else {
		return dictionaryCache
	}
}

// ManagementFactory returns a factory function for creating Management instances.
func ManagementFactory() {
	once.Do(func() {
		config := configReader.Instance()
		cacheType := config.Cache.Type
		if strings.ToUpper(cacheType) == strings.ToUpper("redis") {
			redisCache = newRedisCacheManagement()
			dictionaryCache = nil
		} else {
			dictionaryCache = newDictionaryCacheManagement()
			redisCache = nil
		}
	})
}

// NewDictionaryCacheManagement creates a new instance of DictionaryCacheManager.
func newDictionaryCacheManagement() *DictionaryCacheManager {
	return &DictionaryCacheManager{
		cache: make(map[string]string),
	}
}

// NewRedisCacheManagement creates a new instance of RedisCacheManager.
func newRedisCacheManagement() *RedisCacheManager {
	config := configReader.Instance()
	return &RedisCacheManager{
		client: redis.NewClient(&redis.Options{
			Addr:     config.Cache.Address,
			Username: config.Cache.Username,
			Password: config.Cache.Password,
			DB:       0,
		}),
	}
}
