package config

import (
	"fmt"

	"github.com/ZihxS/be-alibabacloud-genai-2025/constant"
	"github.com/go-redis/redis/v8"
)

func InitRedis(isProduction bool) (redisClient *redis.Client) {
	var redisOptions *redis.Options
	if isProduction {
		redisOptions = &redis.Options{
			Addr:     fmt.Sprintf("%v:%v", constant.REDIS_HOST, constant.REDIS_PORT),
			Password: constant.REDIS_PASS,
		}
	} else {
		redisOptions = &redis.Options{
			Addr: fmt.Sprintf("%v:%v", constant.REDIS_HOST, constant.REDIS_PORT),
		}
	}

	redisClient = redis.NewClient(redisOptions)
	return
}
