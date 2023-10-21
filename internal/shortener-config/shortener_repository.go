package shortenerConfig

type Repository struct {
}

func NewShortenerConfigRepository() *Repository {
	return &Repository{}
}

func (repository *Repository) GetAll() ([]ShortenerConfig, error) {
	return nil, nil
}

func (repository *Repository) GetById(id string) (*ShortenerConfig, error) {
	return nil, nil
}
func (repository *Repository) Create(shortenerConfig ShortenerConfig) (*ShortenerConfig, error) {
	return nil, nil
}
func (repository *Repository) Update(shortenerConfig ShortenerConfig) (*ShortenerConfig, error) {
	return nil, nil
}
func (repository *Repository) Delete(id string) error {
	return nil
}
