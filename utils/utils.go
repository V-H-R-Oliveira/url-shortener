package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"url-shortener/1.0/config"
)

func CreateLogger(prefix string) *log.Logger {
	return log.New(os.Stderr, prefix, log.LstdFlags)
}

func IsMethodAllowed(requestMethod string, allowMethods ...string) bool {
	for _, allowMethod := range allowMethods {
		if requestMethod == allowMethod {
			return true
		}
	}

	return false
}

func GetEnvVariable(key string) (string, error) {
	value, exists := os.LookupEnv(key)

	if exists {
		return value, nil
	}

	return "", fmt.Errorf("key %s not found", key)
}

func IsProd() (bool, error) {
	isProd, err := GetEnvVariable(config.PROD_ENV_VAR)

	if err != nil {
		return false, errors.New("PROD env variable must be set with the values 1 or 0")
	}

	return strconv.ParseBool(isProd)
}

func newTestRequest(path, method string, body map[string]string) *http.Request {
	req := httptest.NewRequest(method, path, nil)

	ctx := context.WithValue(req.Context(),
		config.GetJsonRequestDataContextKey(), body,
	)

	return req.WithContext(ctx)
}

func MakeTestUrlRequest(path, method string, handler http.Handler, body map[string]string) *http.Response {
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, newTestRequest(path, method, body))
	return recorder.Result()
}
