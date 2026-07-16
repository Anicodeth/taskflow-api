package page

import "testing"

func TestClamp(t *testing.T) {
	l, o := Clamp(0, -5, 100)
	if l != 20 || o != 0 {
		t.Fatalf("got l=%d o=%d", l, o)
	}
}

func TestSlice(t *testing.T) {
	got := Slice([]int{1, 2, 3, 4}, 2, 1)
	if len(got) != 2 || got[0] != 2 {
		t.Fatalf("got %v", got)
	}
}