package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestWorkspaceMemoryRoundTrip(t *testing.T) {
	m := NewWorkspaceMemory()
	if _, err := m.CreateWorkspace(models.Workspace{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetWorkspace("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListWorkspace()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestWorkspaceMemoryMissing(t *testing.T) {
	m := NewWorkspaceMemory()
	if _, err := m.GetWorkspace("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}