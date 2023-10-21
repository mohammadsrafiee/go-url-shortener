package shortenerConfigCache

import (
	"encoding/json"
	"fmt"
	shortenerConfigModel "url-shortener/internal/shortener"
	cacheManagement "url-shortener/pkg/cache-management"
	logHandler "url-shortener/pkg/log"
)

var (
	logger = logHandler.Logger()
	cache  = cacheManagement.NewCacheManagerFactory()()
)

type ShortenerConfigCache struct {
}

func NewShortenerConfigCacheRepository() *ShortenerConfigCache {
	return &ShortenerConfigCache{}
}

func (repository *ShortenerConfigCache) GetById(id string) (*shortenerConfigModel.ShortenerConfig, error) {

	var result shortenerConfigModel.ShortenerConfig

	cachedConfig, err := cache.Get(id)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data from cache: %v", err)
	}
	err = json.Unmarshal([]byte(*cachedConfig), &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling data from JSON: %v", err)
	}

	return &result, nil
}

func (repository *ShortenerConfigCache) Create(key string, value *shortenerConfigModel.ShortenerConfig) error {
	jsonData, err := json.Marshal(&value)
	if err != nil {
		logHandler.Logger().Debug("Error marshalling data to JSON:" + err.Error())
		return fmt.Errorf("error marshalling data to JSON: %v", err)
	}

	err = cache.Set(key, jsonData, 0)
	if err != nil {
		logger.Debug("Error marshalling data to JSON:" + err.Error())
		return fmt.Errorf("error marshalling data to JSON: %v", err)
	}
	return nil
}

func (repository *ShortenerConfigCache) Delete(key string) error {
	err := cache.Delete(key)
	if err != nil {
		logger.Debug("Error deleting key:" + err.Error())
		return fmt.Errorf("error retrieving data from cache: %v", err)
	}
	return nil
}
