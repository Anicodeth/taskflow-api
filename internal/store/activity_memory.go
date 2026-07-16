package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// ActivityStore describes persistence operations for activity resources.
type ActivityStore interface {
	CreateActivity(v models.Activity) (models.Activity, error)
	GetActivity(id string) (models.Activity, error)
	ListActivity() ([]models.Activity, error)
}

// ActivityMemory is a thread-safe in-memory ActivityStore.
type ActivityMemory struct {
	mu    sync.RWMutex
	items map[string]models.Activity
}

// NewActivityMemory returns an initialised store.
func NewActivityMemory() *ActivityMemory {
	return &ActivityMemory{items: make(map[string]models.Activity)}
}

func (m *ActivityMemory) CreateActivity(v models.Activity) (models.Activity, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *ActivityMemory) GetActivity(id string) (models.Activity, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Activity{}, ErrNotFound
	}
	return v, nil
}

func (m *ActivityMemory) ListActivity() ([]models.Activity, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Activity, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}