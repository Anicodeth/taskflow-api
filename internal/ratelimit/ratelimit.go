// Package ratelimit provides a simple token-bucket rate limiter.
package ratelimit

import "sync"

// Bucket is a thread-safe token bucket with a fixed capacity.
type Bucket struct {
	mu     sync.Mutex
	tokens int
	cap    int
}

// NewBucket creates a full bucket with the given capacity.
func NewBucket(capacity int) *Bucket {
	return &Bucket{tokens: capacity, cap: capacity}
}

// Allow consumes a token, returning false when the bucket is empty.
func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.tokens <= 0 {
		return false
	}
	b.tokens--
	return true
}

// Refill resets the bucket to full capacity.
func (b *Bucket) Refill() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.tokens = b.cap
}