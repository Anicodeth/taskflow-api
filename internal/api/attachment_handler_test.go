package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/store"
)

func TestAttachmentCreate(t *testing.T) {
	h := NewAttachmentHandler(store.NewAttachmentMemory())
	body := strings.NewReader(`{"name":"demo","description":"d"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/attachments", body)
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}

func TestAttachmentCreateValidation(t *testing.T) {
	h := NewAttachmentHandler(store.NewAttachmentMemory())
	req := httptest.NewRequest(http.MethodPost, "/api/attachments", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d", rec.Code)
	}
}