package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func ConnectRedis() {

	var REDIS_HOST = os.Getenv("REDIS_HOST")
	var REDIS_PORT = os.Getenv("REDIS_PORT")
	var REDIS_USER = os.Getenv("REDIS_USER")
	var REDIS_PASS = os.Getenv("REDIS_PASS")

	redisClient = redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: REDIS_PASS,
		Username: REDIS_USER,
		DB:       0, //Connect to default Redis DB (out of 0–15)
	})
}
