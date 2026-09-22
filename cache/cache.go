package cache

import "time"

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string, ttl time.Duration) error
	Delete(key string) error
}

type Config struct {
	RedisURL string
}

func New(cfg Config) Cache {
	if cfg.RedisURL != "" {
		rc, err := NewRedisCache(cfg.RedisURL)
		if err == nil {
			return rc
		}
	}
	return NewMemoryCache()
}
