package limiter

import (
	"sync"
	"time"
)

type LeakyBucketLimiter struct {
	mu        sync.Mutex
	lastLeak  time.Time
	capacity  int
	remaining int
	leakRate  time.Duration
}

func NewLeakyBucketLimiter(capacity int, leakRate time.Duration) *LeakyBucketLimiter {
	return &LeakyBucketLimiter{
		lastLeak:  time.Now(),
		capacity:  capacity,
		remaining: capacity,
		leakRate:  leakRate,
	}
}

func (l *LeakyBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(l.lastLeak)
	l.lastLeak = now

	leakedRequests := int(elapsedTime / l.leakRate)
	l.remaining += leakedRequests
	if l.remaining > l.capacity {
		l.remaining = l.capacity
	}

	if l.remaining > 0 {
		l.remaining--
		return true
	}

	return false
}
