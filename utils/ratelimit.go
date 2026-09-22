package utils

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	mu       sync.Mutex
	tokens   int
	capacity int
	resetAt  time.Time
	interval time.Duration
}

func newRateLimiter(capacity int, interval time.Duration) *rateLimiter {
	return &rateLimiter{
		tokens:   capacity,
		capacity: capacity,
		resetAt:  time.Now().Add(interval),
		interval: interval,
	}
}

func (rl *rateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetAt) {
		rl.tokens = rl.capacity
		rl.resetAt = now.Add(rl.interval)
	}

	if rl.tokens <= 0 {
		return false
	}

	rl.tokens--
	return true
}

func RateLimit(capacity int, interval time.Duration) gin.HandlerFunc {
	rl := newRateLimiter(capacity, interval)

	return func(c *gin.Context) {
		if !rl.allow() {
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				gin.H{
					"error": "rate limit exceeded",
				},
			)
			return
		}
		c.Next()
	}
}
