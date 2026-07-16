package ptr

import "testing"

func TestOf(t *testing.T) {
	if *Of(5) != 5 {
		t.Fatal("of")
	}
}

func TestOr(t *testing.T) {
	if Or[int](nil, 9) != 9 {
		t.Fatal("nil")
	}
	v := 3
	if Or(&v, 9) != 3 {
		t.Fatal("val")
	}
}