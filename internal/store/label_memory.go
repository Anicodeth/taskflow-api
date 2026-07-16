package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// LabelStore describes persistence operations for label resources.
type LabelStore interface {
	CreateLabel(v models.Label) (models.Label, error)
	GetLabel(id string) (models.Label, error)
	ListLabel() ([]models.Label, error)
}

// LabelMemory is a thread-safe in-memory LabelStore.
type LabelMemory struct {
	mu    sync.RWMutex
	items map[string]models.Label
}

// NewLabelMemory returns an initialised store.
func NewLabelMemory() *LabelMemory {
	return &LabelMemory{items: make(map[string]models.Label)}
}

func (m *LabelMemory) CreateLabel(v models.Label) (models.Label, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *LabelMemory) GetLabel(id string) (models.Label, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Label{}, ErrNotFound
	}
	return v, nil
}

func (m *LabelMemory) ListLabel() ([]models.Label, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Label, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}