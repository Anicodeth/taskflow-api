package retry

import (
	"errors"
	"testing"
)

func TestDoSucceedsEventually(t *testing.T) {
	n := 0
	err := Do(3, func() error {
		n++
		if n < 2 {
			return errors.New("nope")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if n != 2 {
		t.Fatalf("n=%d", n)
	}
}