package slug

import "testing"

func TestMake(t *testing.T) {
	cases := map[string]string{
		"Hello World":  "hello-world",
		"  A/B  C ":    "a-b-c",
		"already-slug": "already-slug",
	}
	for in, want := range cases {
		if got := Make(in); got != want {
			t.Errorf("Make(%q)=%q want %q", in, got, want)
		}
	}
}