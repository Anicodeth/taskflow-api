package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestSprintMemoryRoundTrip(t *testing.T) {
	m := NewSprintMemory()
	if _, err := m.CreateSprint(models.Sprint{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetSprint("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListSprint()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestSprintMemoryMissing(t *testing.T) {
	m := NewSprintMemory()
	if _, err := m.GetSprint("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}