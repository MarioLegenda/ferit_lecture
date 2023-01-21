package redis

import (
	"fmt"
	redisClient "github.com/go-redis/redis"
	"os"
)

var client *redisClient.Client

func NewClient() (*redisClient.Client, error) {
	if client != nil {
		return client, nil
	}

	rdb := redisClient.NewClient(&redisClient.Options{
		Addr:     fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	err := rdb.Set("test", "test", 0).Err()

	if err != nil {
		return nil, err
	}

	rdb.Del("test")

	client = rdb

	return client, nil
}

func Close() error {
	if err := client.Close(); err != nil {
		return err
	}

	return nil
}
