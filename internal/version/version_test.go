package version

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	if !strings.Contains(String(), "taskflow-api") {
		t.Fatalf("got %q", String())
	}
	if Version == "" {
		t.Fatal("version must be set")
	}
}