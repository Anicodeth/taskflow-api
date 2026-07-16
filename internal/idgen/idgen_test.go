package idgen

import "testing"

func TestNewUnique(t *testing.T) {
	seen := map[string]bool{}
	for i := 0; i < 1000; i++ {
		id := New()
		if len(id) != 16 {
			t.Fatalf("bad length: %q", id)
		}
		if seen[id] {
			t.Fatalf("collision: %s", id)
		}
		seen[id] = true
	}
}