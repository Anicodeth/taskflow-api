package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// TagStore describes persistence operations for tag resources.
type TagStore interface {
	CreateTag(v models.Tag) (models.Tag, error)
	GetTag(id string) (models.Tag, error)
	ListTag() ([]models.Tag, error)
}

// TagMemory is a thread-safe in-memory TagStore.
type TagMemory struct {
	mu    sync.RWMutex
	items map[string]models.Tag
}

// NewTagMemory returns an initialised store.
func NewTagMemory() *TagMemory {
	return &TagMemory{items: make(map[string]models.Tag)}
}

func (m *TagMemory) CreateTag(v models.Tag) (models.Tag, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *TagMemory) GetTag(id string) (models.Tag, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Tag{}, ErrNotFound
	}
	return v, nil
}

func (m *TagMemory) ListTag() ([]models.Tag, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Tag, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}