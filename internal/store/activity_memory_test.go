package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestActivityMemoryRoundTrip(t *testing.T) {
	m := NewActivityMemory()
	if _, err := m.CreateActivity(models.Activity{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetActivity("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListActivity()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestActivityMemoryMissing(t *testing.T) {
	m := NewActivityMemory()
	if _, err := m.GetActivity("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}