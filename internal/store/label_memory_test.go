package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestLabelMemoryRoundTrip(t *testing.T) {
	m := NewLabelMemory()
	if _, err := m.CreateLabel(models.Label{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetLabel("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListLabel()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestLabelMemoryMissing(t *testing.T) {
	m := NewLabelMemory()
	if _, err := m.GetLabel("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}