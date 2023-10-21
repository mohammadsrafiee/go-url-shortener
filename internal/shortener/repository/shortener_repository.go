package shortenerConfigRepository

import shortenerConfig "url-shortener/internal/shortener"

type ShortenerConfigRepository struct {
}

func NewShortenerConfigRepository() *ShortenerConfigRepository {
	return &ShortenerConfigRepository{}
}

func (repository *ShortenerConfigRepository) GetAll() ([]shortenerConfig.ShortenerConfig, error) {
	return nil, nil
}

func (repository *ShortenerConfigRepository) GetById(id string) (*shortenerConfig.ShortenerConfig, error) {
	return nil, nil
}
func (repository *ShortenerConfigRepository) Create(shortenerConfig shortenerConfig.ShortenerConfig) (*shortenerConfig.ShortenerConfig, error) {
	return nil, nil
}
func (repository *ShortenerConfigRepository) Update(shortenerConfig shortenerConfig.ShortenerConfig) (*shortenerConfig.ShortenerConfig, error) {
	return nil, nil
}
func (repository *ShortenerConfigRepository) Delete(id string) error {
	return nil
}
