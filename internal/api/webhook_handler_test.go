package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/store"
)

func TestWebhookCreate(t *testing.T) {
	h := NewWebhookHandler(store.NewWebhookMemory())
	body := strings.NewReader(`{"name":"demo","description":"d"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/webhooks", body)
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}

func TestWebhookCreateValidation(t *testing.T) {
	h := NewWebhookHandler(store.NewWebhookMemory())
	req := httptest.NewRequest(http.MethodPost, "/api/webhooks", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d", rec.Code)
	}
}