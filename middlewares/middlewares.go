package middlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"url-shortener/1.0/config"
	"url-shortener/1.0/utils"
)

func ValidateRequestMethodMiddleware(next http.Handler, methods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if utils.IsMethodAllowed(r.Method, methods...) {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Invalid method", http.StatusBadRequest)
	}
}

func DecodeJsonResponseMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		decoder := json.NewDecoder(r.Body)

		var data map[string]string

		if err := decoder.Decode(&data); err != nil {
			log.Println("Failed to parse response due error:", err)
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), config.GetJsonRequestDataContextKey(), data)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
