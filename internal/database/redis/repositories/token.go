package repositories

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepo(redis *redis.Client) *RedisRepository {
	return &RedisRepository{
		redis: redis,
	}
}

func (t RedisRepository) Get(s string) (string, error) {
	val, err := t.redis.Get(context.Background(), s).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

func (t RedisRepository) Set(key string, value string, expiration time.Duration) error {
	err := t.redis.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (t RedisRepository) Delete(s string) error {
	_, err := t.redis.Del(context.Background(), s).Result()

	if err != nil {
		return err
	}

	return nil
}
