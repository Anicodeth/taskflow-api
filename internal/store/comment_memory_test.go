package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestCommentMemoryRoundTrip(t *testing.T) {
	m := NewCommentMemory()
	if _, err := m.CreateComment(models.Comment{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetComment("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListComment()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestCommentMemoryMissing(t *testing.T) {
	m := NewCommentMemory()
	if _, err := m.GetComment("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}