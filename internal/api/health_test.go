package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/store"
)

func TestHealthz(t *testing.T) {
	srv := NewRouter(store.NewMemory())
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	srv.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}