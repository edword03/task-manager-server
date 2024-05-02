package redis

import (
	"github.com/redis/go-redis/v9"
	"task-manager/internal/infrastructure/config"
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
}
