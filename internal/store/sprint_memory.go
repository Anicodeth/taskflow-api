package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// SprintStore describes persistence operations for sprint resources.
type SprintStore interface {
	CreateSprint(v models.Sprint) (models.Sprint, error)
	GetSprint(id string) (models.Sprint, error)
	ListSprint() ([]models.Sprint, error)
}

// SprintMemory is a thread-safe in-memory SprintStore.
type SprintMemory struct {
	mu    sync.RWMutex
	items map[string]models.Sprint
}

// NewSprintMemory returns an initialised store.
func NewSprintMemory() *SprintMemory {
	return &SprintMemory{items: make(map[string]models.Sprint)}
}

func (m *SprintMemory) CreateSprint(v models.Sprint) (models.Sprint, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *SprintMemory) GetSprint(id string) (models.Sprint, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Sprint{}, ErrNotFound
	}
	return v, nil
}

func (m *SprintMemory) ListSprint() ([]models.Sprint, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Sprint, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}