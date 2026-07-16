package validate

import "testing"

func TestRequired(t *testing.T) {
	if err := Required("name", ""); err == nil {
		t.Fatal("expected error")
	}
	if err := Required("name", "x"); err != nil {
		t.Fatalf("unexpected: %v", err)
	}
}

func TestMaxLen(t *testing.T) {
	if err := MaxLen("name", "abcd", 3); err == nil {
		t.Fatal("expected error")
	}
}