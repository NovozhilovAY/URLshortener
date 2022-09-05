package repository

type Repository interface {
	InsertUrl(code string, originalURL string)
	GetUrl(code string) string
	ContainsUrl(originalURL string) (bool, string)
}
