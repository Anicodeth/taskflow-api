package problem

import (
	"net/http"
	"testing"
)

func TestProblemError(t *testing.T) {
	p := New("not_found", http.StatusNotFound, "missing")
	if p.Error() != "not_found: missing" {
		t.Fatalf("got %q", p.Error())
	}
}