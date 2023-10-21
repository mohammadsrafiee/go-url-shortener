package shortenerConfigService

import (
	shortenerConfig "url-shortener/internal/shortener"
	shortenerConfigCache "url-shortener/internal/shortener/cache"
	shortenerConfigRepository "url-shortener/internal/shortener/repository"
	cacheManagement "url-shortener/pkg/cache-management"
	logHandler "url-shortener/pkg/log"
	_ "url-shortener/pkg/shorter"
)

var (
	cache           = cacheManagement.NewCacheManagerFactory()()
	logger          = logHandler.Logger()
	repository      = shortenerConfigRepository.NewShortenerConfigRepository()
	cacheRepository = shortenerConfigCache.NewShortenerConfigCacheRepository()
)

type StorageService interface {
	GetAll() ([]shortenerConfig.ShortenerConfig, error)
	GetByID(id int) (*shortenerConfig.ShortenerConfig, error)
	Create(shortenerConfig shortenerConfig.ShortenerConfig) (*shortenerConfig.ShortenerConfig, error)
	Update(shortenerConfig shortenerConfig.ShortenerConfig) (*shortenerConfig.ShortenerConfig, error)
	Delete(id string) error
}

type ShortenerConfigService struct {
}

func NewShortenerConfigService() *ShortenerConfigService {
	return &ShortenerConfigService{}
}

func (service *ShortenerConfigService) GetAll() ([]shortenerConfig.ShortenerConfig, error) {
	result, err := repository.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *ShortenerConfigService) GetByID(id string) (*shortenerConfig.ShortenerConfig, error) {
	result, err := cacheRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	if result != nil {
		return result, nil
	}

	result, err = repository.GetById(id)
	if err != nil {
		return nil, err
	}

	if result != nil {
		err = cacheRepository.Create(result.ID, result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (service *ShortenerConfigService) Create(data shortenerConfig.ShortenerConfig) (*shortenerConfig.ShortenerConfig, error) {

	result, err := repository.Create(data)
	if err != nil {
		return nil, err
	}
	err = cacheRepository.Create(data.ID, &data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *ShortenerConfigService) Update(data shortenerConfig.ShortenerConfig) (*shortenerConfig.ShortenerConfig, error) {
	err := cacheRepository.Delete(data.ID)
	if err != nil {
		return nil, err
	}

	result, err := repository.Update(data)
	if err != nil {
		return nil, err
	}

	err = cacheRepository.Create(result.ID, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *ShortenerConfigService) Delete(id string) error {
	err := cacheRepository.Delete(id)
	if err != nil {
		return err
	}

	err = repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
