package shortener

import (
	goUrl "net/url"
	"testing"
	"url-shortener/1.0/config"
	urlencoder "url-shortener/1.0/url-encoder"
)

func TestUrlShortener(t *testing.T) {
	t.Run("Test invalid urls", func(t *testing.T) {
		tests := []struct {
			name    string
			url     string
			wantErr bool
		}{
			{
				name:    "relative url",
				url:     "/abc",
				wantErr: true,
			},
			{
				name:    "invalid url",
				url:     "abc",
				wantErr: true,
			},
			{
				name:    "empty url",
				url:     "",
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := NewUrlShortener(tt.url).ValidateUrl(); (err != nil) != tt.wantErr {
					t.Errorf("UrlShortener.ValidateUrl() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})

	t.Run("Test valid urls", func(t *testing.T) {
		tests := []struct {
			name    string
			url     string
			wantErr bool
		}{
			{
				name:    "absolute url without www",
				url:     "https://google.com",
				wantErr: false,
			},
			{
				name:    "absolute url with final /",
				url:     "https://www.test.com/",
				wantErr: false,
			},
			{
				name:    "absolute url with path and query string",
				url:     "https://www.test.com/abc?q=123",
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := NewUrlShortener(tt.url).ValidateUrl(); (err != nil) != tt.wantErr {
					t.Errorf("UrlShortener.ValidateUrl() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})

	t.Run("Test short url format", func(t *testing.T) {
		tests := []struct {
			name string
			url  string
		}{
			{
				name: "absolute url without www",
				url:  "https://google.com",
			},
			{
				name: "absolute url with final /",
				url:  "https://www.test.com/",
			},
			{
				name: "absolute url with path and query string",
				url:  "https://www.test.com/abc?q=123",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				shortUrl, err := goUrl.Parse(NewUrlShortener(tt.url).ShortUrl())

				if err != nil {
					t.Fatal("Invalid url format")
				}

				if config.BASE_URL.Hostname() != shortUrl.Hostname() {
					t.Fatalf("Expected hostname %s, got hostname %s\n", config.BASE_URL.Hostname(), shortUrl.Hostname())
				}

				if urlPath := urlencoder.GetEncodedUrlPath(shortUrl.String()); urlPath == "" {
					t.Fatal("Missing base62 hash")
				}
			})
		}
	})
}
