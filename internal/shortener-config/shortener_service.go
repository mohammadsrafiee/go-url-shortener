package shortenerConfig

import (
	_ "url-shortener/pkg/shorter"
)

var (
	//cache           = cacheManagement.NewCacheManagerFactory()()
	//logger          = logHandler.Logger()
	repository      *Repository
	cacheRepository *Cache
)

type StorageService interface {
	GetAll() ([]ShortenerConfig, error)
	GetByID(id int) (*ShortenerConfig, error)
	Create(shortenerConfig ShortenerConfig) (*ShortenerConfig, error)
	Update(shortenerConfig ShortenerConfig) (*ShortenerConfig, error)
	Delete(id string) error
}

type Service struct {
}

func NewShortenerConfigService() *Service {
	repository = NewShortenerConfigRepository()
	cacheRepository = NewShortenerConfigCacheRepository()
	return &Service{}
}

func (service *Service) GetAll() ([]ShortenerConfig, error) {
	result, err := repository.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *Service) GetByID(id string) (*ShortenerConfig, error) {
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

func (service *Service) Create(data ShortenerConfig) (*ShortenerConfig, error) {

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

func (service *Service) Update(data ShortenerConfig) (*ShortenerConfig, error) {
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

func (service *Service) Delete(id string) error {
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
