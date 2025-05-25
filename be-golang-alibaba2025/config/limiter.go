package config

import (
	"context"
	"fmt"
	"log"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/ZihxS/be-alibabacloud-genai-2025/constant"
	"github.com/redis/go-redis/v9"
)

func InitLimiterStorage(isProduction bool) ratelimit.Store {
	var redisClient *redis.Client

	if isProduction {
		redisOptions := &redis.Options{
			Addr:     fmt.Sprintf("%v:%v", constant.REDIS_HOST, constant.REDIS_PORT),
			Password: constant.REDIS_PASS,
		}

		redisClient = redis.NewClient(redisOptions)
	} else {
		redisOptions := &redis.Options{
			Addr: fmt.Sprintf("%v:%v", constant.REDIS_HOST, constant.REDIS_PORT),
		}

		redisClient = redis.NewClient(redisOptions)
	}

	redisCtx := context.Background()
	if _, err := redisClient.Ping(redisCtx).Result(); err != nil {
		if err.Error() == constant.REDIS_ERR_NO_PASSWORD {
			redisOptions := &redis.Options{
				Addr: fmt.Sprintf("%v:%v", constant.REDIS_HOST, constant.REDIS_PORT),
			}

			redisClient = redis.NewClient(redisOptions)

			if _, err := redisClient.Ping(redisCtx).Result(); err != nil {
				log.Fatal("error connection to redis server, error: ", err.Error())
			}
		} else {
			log.Fatal("error connection to redis server, error: ", err.Error())
		}
	}

	return ratelimit.RedisStore(&ratelimit.RedisOptions{
		RedisClient: redisClient,
		Rate:        time.Second,
		Limit:       20, // TODO: please using env
	})
}
