package caching

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var rds *redis.Client

func Init() {
	rds = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
}

func GetRedisClient() *redis.Client {
	return rds
}
