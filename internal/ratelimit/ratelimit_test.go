package ratelimit

import "testing"

func TestBucket(t *testing.T) {
	b := NewBucket(2)
	if !b.Allow() || !b.Allow() {
		t.Fatal("should allow 2")
	}
	if b.Allow() {
		t.Fatal("should be empty")
	}
	b.Refill()
	if !b.Allow() {
		t.Fatal("refill")
	}
}