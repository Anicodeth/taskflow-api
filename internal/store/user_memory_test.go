package store

import (
	"testing"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

func TestUserMemoryRoundTrip(t *testing.T) {
	m := NewUserMemory()
	if _, err := m.CreateUser(models.User{ID: "1", Name: "x"}); err != nil {
		t.Fatalf("create: %v", err)
	}
	got, err := m.GetUser("1")
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Name != "x" {
		t.Fatalf("want x, got %q", got.Name)
	}
	all, err := m.ListUser()
	if err != nil || len(all) != 1 {
		t.Fatalf("list: %v len=%d", err, len(all))
	}
}

func TestUserMemoryMissing(t *testing.T) {
	m := NewUserMemory()
	if _, err := m.GetUser("nope"); err == nil {
		t.Fatal("expected error for missing record")
	}
}