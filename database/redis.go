package database

import (
	"context"
	"log"
	"time"

	"url-shortener/1.0/config"
	urlencoder "url-shortener/1.0/url-encoder"
	"url-shortener/1.0/utils"

	goRedis "github.com/go-redis/redis/v9"
)

type RedisDatabase struct {
	client *goRedis.Client
}

func createRedisContextOperation() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

func NewRedisDatabase() Repository {
	redisAddr, err := utils.GetEnvVariable(config.REDIS_URL_ENV_VAR)

	if err != nil {
		log.Fatalf("%s not defined\n", config.REDIS_URL_ENV_VAR)
	}

	return &RedisDatabase{
		client: goRedis.NewClient(&goRedis.Options{
			Addr: redisAddr,
		}),
	}
}

func (redis *RedisDatabase) AddUrl(url, shortUrl string) error {
	ctx, cancel := createRedisContextOperation()
	defer cancel()
	return redis.client.Set(ctx, urlencoder.GetEncodedUrlPath(shortUrl), url, config.ONE_MONTH_TTL).Err()
}

func (redis *RedisDatabase) GetUrl(shortUrl string) (string, error) {
	ctx, cancel := createRedisContextOperation()
	defer cancel()
	return redis.client.Get(ctx, urlencoder.GetEncodedUrlPath(shortUrl)).Result()
}
