package config

import (
	"os"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost" // fallback default, harusnya "redis" saat pakai docker-compose
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	RedisClient = redis.NewClient(&redis.Options{
		Addr: addr,
		// Password: "", // kalau pakai password, isi dari os.Getenv("REDIS_PASS")
		DB: 0,
	})
}
