package envx

import (
	"os"
	"testing"
)

func TestString(t *testing.T) {
	os.Setenv("ENVX_T", "hi")
	if String("ENVX_T", "x") != "hi" {
		t.Fatal("string")
	}
	if String("ENVX_MISSING", "x") != "x" {
		t.Fatal("fallback")
	}
}

func TestInt(t *testing.T) {
	os.Setenv("ENVX_N", "42")
	if Int("ENVX_N", 0) != 42 {
		t.Fatal("int")
	}
}