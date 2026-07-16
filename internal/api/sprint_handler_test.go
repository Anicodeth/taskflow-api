package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/store"
)

func TestSprintCreate(t *testing.T) {
	h := NewSprintHandler(store.NewSprintMemory())
	body := strings.NewReader(`{"name":"demo","description":"d"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/sprints", body)
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}

func TestSprintCreateValidation(t *testing.T) {
	h := NewSprintHandler(store.NewSprintMemory())
	req := httptest.NewRequest(http.MethodPost, "/api/sprints", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d", rec.Code)
	}
}