package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestAttachmentMemoryRoundTrip(t *testing.T) {
	m := NewAttachmentMemory()
	if _, err := m.CreateAttachment(models.Attachment{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetAttachment("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListAttachment()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestAttachmentMemoryMissing(t *testing.T) {
	m := NewAttachmentMemory()
	if _, err := m.GetAttachment("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}