package config

import (
	"net/url"
	"time"
)

const ONE_MONTH_TTL = time.Hour * 730
const JSON_REQUEST_DATA_KEY = "jsonRequestData"
const REDIS_URL_ENV_VAR = "REDIS_URL"
const PROD_ENV_VAR = "PROD"
const DEFAULT_SERVER_PORT = "8080"

const SERVER_READ_TIMEOUT = 2 * time.Second
const SERVER_WRITE_TIMEOUT = 2 * time.Second

var BASE_URL, _ = url.Parse("http://shortener.ly")

type middlewareContextKey string

func getMiddlewareContextKey(label string) middlewareContextKey {
	return middlewareContextKey(label)
}

func GetJsonRequestDataContextKey() middlewareContextKey {
	return getMiddlewareContextKey(JSON_REQUEST_DATA_KEY)
}
