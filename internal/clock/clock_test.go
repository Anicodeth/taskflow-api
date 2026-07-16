package clock

import (
	"testing"
	"time"
)

func TestSystemUTC(t *testing.T) {
	now := System()
	if now.Location() != time.UTC {
		t.Fatalf("expected UTC, got %v", now.Location())
	}
}