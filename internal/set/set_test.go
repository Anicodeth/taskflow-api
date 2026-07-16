package set

import "testing"

func TestSet(t *testing.T) {
	s := New(1, 2, 2, 3)
	if s.Len() != 3 {
		t.Fatalf("len=%d", s.Len())
	}
	if !s.Has(2) {
		t.Fatal("has")
	}
	s.Add(4)
	if !s.Has(4) {
		t.Fatal("add")
	}
}