package controllers

import (
	"fmt"
	"net/http"
	"url-shortener/1.0/config"
	"url-shortener/1.0/database"
	"url-shortener/1.0/shortener"
	"url-shortener/1.0/utils"
)

func ShortenerUrl(repository database.Repository) http.HandlerFunc {
	logger := utils.CreateLogger("[shortener-controller]> ")

	return func(w http.ResponseWriter, r *http.Request) {
		dataMap := r.Context().Value(config.GetJsonRequestDataContextKey()).(map[string]string)
		url, hasUrl := dataMap["url"]

		if !hasUrl {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		shortener := shortener.NewUrlShortener(url)

		if err := shortener.ValidateUrl(); err != nil {
			http.Error(w, "Invalid url", http.StatusBadRequest)
			return
		}

		shortUrl := shortener.ShortUrl()

		if err := repository.AddUrl(shortener.Url, shortUrl); err != nil {
			logger.Println("Failed to save the url due error", err)
			http.Error(w, "Short url generation failed", http.StatusInternalServerError)
		}

		fmt.Fprintln(w, shortUrl)
	}
}

func RedirectUrl(repository database.Repository) http.HandlerFunc {
	logger := utils.CreateLogger("[redirect-url-controller]> ")

	return func(w http.ResponseWriter, r *http.Request) {
		if originalUrl, err := repository.GetUrl(r.URL.EscapedPath()); err != nil {
			logger.Println("Failed to fetch the original url due error:", err)
			http.Error(w, "url not found", http.StatusNotFound)
		} else {
			http.Redirect(w, r, originalUrl, http.StatusPermanentRedirect)
		}
	}
}
