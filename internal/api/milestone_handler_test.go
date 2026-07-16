package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/store"
)

func TestMilestoneCreate(t *testing.T) {
	h := NewMilestoneHandler(store.NewMilestoneMemory())
	body := strings.NewReader(`{"name":"demo","description":"d"}`)
	req := httptest.NewRequest(http.MethodPost, "/api/milestones", body)
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}

func TestMilestoneCreateValidation(t *testing.T) {
	h := NewMilestoneHandler(store.NewMilestoneMemory())
	req := httptest.NewRequest(http.MethodPost, "/api/milestones", strings.NewReader(`{}`))
	rec := httptest.NewRecorder()
	h.Create(rec, req)
	if rec.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d", rec.Code)
	}
}