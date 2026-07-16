package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestWebhookMemoryRoundTrip(t *testing.T) {
	m := NewWebhookMemory()
	if _, err := m.CreateWebhook(models.Webhook{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetWebhook("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListWebhook()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestWebhookMemoryMissing(t *testing.T) {
	m := NewWebhookMemory()
	if _, err := m.GetWebhook("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}