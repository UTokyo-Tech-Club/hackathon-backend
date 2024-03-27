package redis

import (
	"context"
	"hackathon-backend/utils/logger"
	"os"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis() {
	logger := logger.NewLogger()

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		logger.Fatal("REDIS_ADDR environment variable not set")
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		logger.Fatal("Error connecting to redis: ", err)
	}

	logger.Info("Connected to redis")
}

func IncrementRedis() {
	logger := logger.NewLogger()

	ctx := context.Background()

	// Increment visit count.
	err := redisClient.Incr(ctx, "visits").Err()
	if err != nil {
		logger.Error("Error incrementing visit count: ", err)
		return
	}

	// Retrieve visit count.
	visits, err := redisClient.Get(ctx, "visits").Result()
	if err != nil {
		logger.Error("Error retrieving visit count: ", err)
		return
	}

	logger.Info("Visit count: ", visits)
}
