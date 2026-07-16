package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestProjectMemoryRoundTrip(t *testing.T) {
	m := NewProjectMemory()
	if _, err := m.CreateProject(models.Project{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetProject("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListProject()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestProjectMemoryMissing(t *testing.T) {
	m := NewProjectMemory()
	if _, err := m.GetProject("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}