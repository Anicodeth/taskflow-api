package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/store"
)

func TestUserCreate(t *testing.T) {
	h := NewUserHandler(store.NewUserMemory())
	body := strings.NewReader(`{"name":"demo","description":"d"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/users", body)
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}

func TestUserCreateValidation(t *testing.T) {
	h := NewUserHandler(store.NewUserMemory())
	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d", rec.Code)
	}
}