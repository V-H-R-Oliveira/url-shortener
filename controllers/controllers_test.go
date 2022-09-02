package controllers

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"url-shortener/1.0/database"
	"url-shortener/1.0/utils"
)

func parseTestResponseBody(body io.ReadCloser) (string, error) {
	defer body.Close()
	response, err := bufio.NewReader(body).ReadString('\n')
	return strings.TrimSpace(response), err
}

func makeShortenerUrlRequest(repository database.Repository, body map[string]string) *http.Response {
	return utils.MakeTestUrlRequest("/shortener", http.MethodPost, ShortenerUrl(repository), body)
}

func makeRedirectUrlRequest(shortUrl string, repository database.Repository) *http.Response {
	return utils.MakeTestUrlRequest(shortUrl, http.MethodGet, RedirectUrl(repository), nil)
}

func getShortUrl(repository database.Repository, url string) string {
	shortUrlResponse := makeShortenerUrlRequest(repository, map[string]string{"url": url})
	shortUrl, _ := parseTestResponseBody(shortUrlResponse.Body)
	return shortUrl
}

func TestShortenerUrl(t *testing.T) {
	t.Run("Should receive a valid url in the response", func(t *testing.T) {
		response := makeShortenerUrlRequest(database.NewMemoryDatabase(), map[string]string{"url": "https://google.com"})
		shortUrl, err := parseTestResponseBody(response.Body)

		if err != nil || shortUrl == "" {
			t.Fatal(err)
		}

		if _, err := url.Parse(shortUrl); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Should return a bad request status if the body is malformed", func(t *testing.T) {
		response := makeShortenerUrlRequest(database.NewMemoryDatabase(), map[string]string{"abc": "https://google.com"})

		if response.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected response status %d, got reponse status %d\n", http.StatusBadRequest, response.StatusCode)
		}
	})

	t.Run("Should return a bad request status if the body is empty", func(t *testing.T) {
		response := makeShortenerUrlRequest(database.NewMemoryDatabase(), nil)

		if response.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected response status %d, got reponse status %d\n", http.StatusBadRequest, response.StatusCode)
		}
	})

	t.Run("Should return a bad request status if the requested url is invalid", func(t *testing.T) {
		response := makeShortenerUrlRequest(database.NewMemoryDatabase(), map[string]string{"url": "abc"})

		if response.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected response status %d, got reponse status %d\n", http.StatusBadRequest, response.StatusCode)
		}
	})

	t.Run("Should return a bad request status if the requested url is relative", func(t *testing.T) {
		response := makeShortenerUrlRequest(database.NewMemoryDatabase(), map[string]string{"url": "/abc"})

		if response.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected response status %d, got reponse status %d\n", http.StatusBadRequest, response.StatusCode)
		}
	})

	t.Run("Should return an internal server error status if it fails to save the url", func(t *testing.T) {
		response := makeShortenerUrlRequest(database.NewMockFailingOperationsDatabase(), map[string]string{"url": "https://google.com"})

		if response.StatusCode != http.StatusInternalServerError {
			t.Fatalf("expected response status %d, got reponse status %d\n", http.StatusInternalServerError, response.StatusCode)
		}
	})
}

func TestRedirectUrl(t *testing.T) {

	t.Run("Should return a permanent redirect status if the redirect was successfully made", func(t *testing.T) {
		const url = "https://google.com"
		testDatabase := database.NewMemoryDatabase()
		shortUrl := getShortUrl(testDatabase, url)
		response := makeRedirectUrlRequest(shortUrl, testDatabase)

		if response.StatusCode != http.StatusPermanentRedirect {
			t.Fatalf("Expected response status %d, got %d\n", http.StatusPermanentRedirect, response.StatusCode)
		}
	})

	t.Run("Should redirect to the correct url", func(t *testing.T) {
		const url = "https://google.com"
		testDatabase := database.NewMemoryDatabase()
		shortUrl := getShortUrl(testDatabase, url)
		response := makeRedirectUrlRequest(shortUrl, testDatabase)

		location, err := response.Location()

		if err != nil {
			t.Fatal(err)
		}

		if location.String() != url {
			t.Fatalf("Expected redirect to %s, got redirect to %s\n", url, location.String())
		}
	})

	t.Run("Should return not found status if the requested url was not found", func(t *testing.T) {
		response := makeRedirectUrlRequest("http://shortener.ly/abc", database.NewMemoryDatabase())

		if response.StatusCode != http.StatusNotFound {
			t.Fatalf("Expected response status code %d, got response status code %d\n", http.StatusNotFound, response.StatusCode)
		}
	})
}
