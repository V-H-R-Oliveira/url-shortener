package database

import "errors"

type mockFailingOperationsDatabase struct {
}

func NewMockFailingOperationsDatabase() Repository {
	return &mockFailingOperationsDatabase{}
}

func (db *mockFailingOperationsDatabase) AddUrl(url, shortUrlHash string) error {
	return errors.New("failed to add url")
}

func (db *mockFailingOperationsDatabase) GetUrl(shortUrlHash string) (string, error) {
	return "", errors.New("failed to get the url")
}
