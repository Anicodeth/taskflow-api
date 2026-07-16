package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestNotificationMemoryRoundTrip(t *testing.T) {
	m := NewNotificationMemory()
	if _, err := m.CreateNotification(models.Notification{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetNotification("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListNotification()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestNotificationMemoryMissing(t *testing.T) {
	m := NewNotificationMemory()
	if _, err := m.GetNotification("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}