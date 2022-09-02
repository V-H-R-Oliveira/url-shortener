package shortener

import (
	"errors"
	goUrl "net/url"
	"url-shortener/1.0/config"
	urlencoder "url-shortener/1.0/url-encoder"
)

type UrlShortener struct {
	Url string
}

func NewUrlShortener(url string) *UrlShortener {
	return &UrlShortener{
		Url: url,
	}
}

func (shortener *UrlShortener) saveEscapedUrl(escapedUrl string) {
	shortener.Url = escapedUrl
}

func (shortner *UrlShortener) ValidateUrl() error {
	parsedUrl, err := goUrl.ParseRequestURI(shortner.Url)

	if err != nil {
		return err
	}

	if parsedUrl.IsAbs() {
		shortner.saveEscapedUrl(parsedUrl.String())
		return nil
	}

	return errors.New("relative urls are not allowed")
}

func (shortener *UrlShortener) formatShortUrl(urlPath string) string {
	shortUrl := &goUrl.URL{
		Scheme: config.BASE_URL.Scheme,
		Host:   config.BASE_URL.Host,
		Path:   urlPath,
	}

	return shortUrl.String()
}

func (shortner *UrlShortener) ShortUrl() string {
	return shortner.formatShortUrl(urlencoder.EncodeUrlPath(shortner.Url))
}
