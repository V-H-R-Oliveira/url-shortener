package database

import (
	"errors"
	urlencoder "url-shortener/1.0/url-encoder"
)

type memoryDatabase struct {
	database map[string]string
}

func NewMemoryDatabase() Repository {
	return &memoryDatabase{
		database: make(map[string]string),
	}
}

func (db *memoryDatabase) AddUrl(url, shortUrl string) error {
	db.database[urlencoder.GetEncodedUrlPath(shortUrl)] = url
	return nil
}

func (db *memoryDatabase) GetUrl(shortUrl string) (string, error) {
	if url, hasUrl := db.database[urlencoder.GetEncodedUrlPath(shortUrl)]; hasUrl {
		return url, nil
	}

	return "", errors.New("url not found")
}
