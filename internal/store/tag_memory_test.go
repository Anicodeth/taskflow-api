package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestTagMemoryRoundTrip(t *testing.T) {
	m := NewTagMemory()
	if _, err := m.CreateTag(models.Tag{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetTag("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListTag()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestTagMemoryMissing(t *testing.T) {
	m := NewTagMemory()
	if _, err := m.GetTag("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}