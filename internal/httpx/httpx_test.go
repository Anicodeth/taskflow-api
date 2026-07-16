package httpx

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"a":1}`))
	var out struct {
		A int `json:"a"`
	}
	if err := DecodeJSON(req, &out); err != nil {
		t.Fatal(err)
	}
	if out.A != 1 {
		t.Fatalf("got %d", out.A)
	}
}