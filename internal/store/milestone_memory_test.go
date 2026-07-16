package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestMilestoneMemoryRoundTrip(t *testing.T) {
	m := NewMilestoneMemory()
	if _, err := m.CreateMilestone(models.Milestone{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetMilestone("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListMilestone()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestMilestoneMemoryMissing(t *testing.T) {
	m := NewMilestoneMemory()
	if _, err := m.GetMilestone("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}