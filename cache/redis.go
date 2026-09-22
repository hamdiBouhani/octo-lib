package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(url string) (Cache, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return &RedisCache{
		client: redis.NewClient(opt),
	}, nil
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.client.Get(context.Background(), key).Result()
}

func (c *RedisCache) Set(key, value string, ttl time.Duration) error {
	return c.client.Set(context.Background(), key, value, ttl).Err()
}

func (c *RedisCache) Delete(key string) error {
	return c.client.Del(context.Background(), key).Err()
}
