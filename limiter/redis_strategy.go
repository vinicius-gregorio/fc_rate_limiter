package limiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRateLimiter struct {
	Client *redis.Client
}

func NewRedisRateLimiter(addr, password string) *RedisRateLimiter {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	return &RedisRateLimiter{Client: client}
}

func (r *RedisRateLimiter) Allow(key string, limit int, duration int64) (bool, error) {
	ctx := context.Background()

	count, err := r.Client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		r.Client.Expire(ctx, key, time.Duration(duration)*time.Second)
	}

	if count > int64(limit) {
		return false, nil
	}
	return true, nil
}

func (r *RedisRateLimiter) Block(key string, duration int64) error {
	ctx := context.Background()
	return r.Client.Set(ctx, key, -1, time.Duration(duration)*time.Second).Err()
}
