package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"task-manager/internal/config"
)

var CacheClient *redis.Client
var TokensClient *redis.Client

func InitRedis(cfg *config.DBConfig) {
	CacheClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: cfg.RedisPass,
		DB:       0,
	})

	TokensClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: cfg.RedisPass,
		DB:       1,
	})

	if TokensClient == nil {
		logrus.Fatalf("init redis fail")
	}

	err := TokensClient.Ping(context.Background()).Err()
	if err != nil {
		logrus.Fatalf("init redis fail")
	}

	TokensClient.AddHook(loggerHook{})
}
