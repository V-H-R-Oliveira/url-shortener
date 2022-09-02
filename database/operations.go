package database

type Repository interface {
	AddUrl(url, shortUrl string) error
	GetUrl(shortUrl string) (string, error)
}
