package config

import (
	"os"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() *redis.Client {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	return redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
}
