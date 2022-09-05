package repository

type InMemoryRepository struct {
	pairsURL map[string]string
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{map[string]string{}}
}

func (i *InMemoryRepository) GetUrl(code string) string {
	var originalURL string = i.pairsURL[code]
	return originalURL
}

func (i *InMemoryRepository) InsertUrl(code string, originalURL string) {
	i.pairsURL[code] = originalURL
}

func (i *InMemoryRepository) ContainsUrl(originalUrl string) (bool, string) {
	for key, value := range i.pairsURL {
		if value == originalUrl {
			return true, key
		}
	}
	return false, ""
}
