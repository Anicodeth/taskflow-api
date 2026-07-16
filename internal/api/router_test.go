package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/store"
)

func TestUnknownRoute(t *testing.T) {
	srv := NewRouter(store.NewMemory())
	req := httptest.NewRequest(http.MethodGet, "/nope", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", rec.Code)
	}
}