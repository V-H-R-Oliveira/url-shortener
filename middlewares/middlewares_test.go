package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-shortener/1.0/utils"
)

func TestValidateRequestMethodMiddleware(t *testing.T) {
	var mockHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }

	t.Run("Should return a bad request status code if the method is not allowed", func(t *testing.T) {

		response := utils.MakeTestUrlRequest(
			"/",
			http.MethodDelete,
			ValidateRequestMethodMiddleware(mockHandler, http.MethodGet, http.MethodPost),
			nil,
		)

		if response.StatusCode != http.StatusBadRequest {
			t.Fatalf("Expected status code 400, got status code %d\n", response.StatusCode)
		}
	})

	t.Run("Should not return a bad request status code if the method is allowed", func(t *testing.T) {
		response := utils.MakeTestUrlRequest(
			"/",
			http.MethodPost,
			ValidateRequestMethodMiddleware(mockHandler, http.MethodGet, http.MethodPost),
			nil,
		)

		if response.StatusCode == http.StatusBadRequest {
			t.Fatal("Should not return a bad request status code\n")
		}
	})
}

func TestDecodeJsonResponseMiddleware(t *testing.T) {
	var mockHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }

	makeRequest := func(request *http.Request) *http.Response {
		response := httptest.NewRecorder()
		DecodeJsonResponseMiddleware(mockHandler)(response, request)
		return response.Result()
	}

	encodeBody := func(body interface{}) io.Reader {
		reader := &bytes.Buffer{}
		json.NewEncoder(reader).Encode(body)
		return reader
	}

	t.Run("Should return a bad request status code if the method is not allowed", func(t *testing.T) {
		response := makeRequest(httptest.NewRequest(http.MethodPost, "/", encodeBody(map[string]string{"url": "https://google.com"})))

		if response.StatusCode != http.StatusOK {
			t.Fatalf("Expected reponse status code 200, got response status code %d\n", response.StatusCode)
		}
	})

	t.Run("Should return a bad request status code if the method is not allowed", func(t *testing.T) {
		response := makeRequest(httptest.NewRequest(http.MethodPost, "/", encodeBody("xyz")))

		if response.StatusCode != http.StatusBadRequest {
			t.Fatalf("Expected reponse status code 400, got response status code %d\n", response.StatusCode)
		}
	})
}
