package repositories

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type ITokenRepository interface {
	Get(string) (string, error)
	Set(key string, value string, expiration time.Duration) error
	Delete(string) error
}

type TokenRepository struct {
	redis *redis.Client
}

func NewTokenRepo(redis *redis.Client) *TokenRepository {
	return &TokenRepository{
		redis: redis,
	}
}

func (t TokenRepository) Get(s string) (string, error) {
	val, err := t.redis.Get(context.Background(), s).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

func (t TokenRepository) Set(key string, value string, expiration time.Duration) error {
	err := t.redis.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (t TokenRepository) Delete(s string) error {
	_, err := t.redis.Del(context.Background(), s).Result()

	if err != nil {
		return err
	}

	return nil
}
