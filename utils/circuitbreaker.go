package utils

import (
	"errors"
	"sync"
	"time"
)

var ErrCircuitOpen = errors.New("circuit open")

type CircuitBreaker struct {
	mu          sync.Mutex
	failures    int
	threshold   int
	openUntil   time.Time
	openTimeout time.Duration
}

func NewCircuitBreaker(threshold int, openTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		threshold:   threshold,
		openTimeout: openTimeout,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if time.Now().Before(cb.openUntil) {
		return ErrCircuitOpen
	}

	err := fn()
	if err != nil {
		cb.failures++
		if cb.failures >= cb.threshold {
			cb.openUntil = time.Now().Add(cb.openTimeout)
		}
		return err
	}

	cb.failures = 0
	return nil
}
