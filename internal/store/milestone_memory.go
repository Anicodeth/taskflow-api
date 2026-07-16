package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// MilestoneStore describes persistence operations for milestone resources.
type MilestoneStore interface {
	CreateMilestone(v models.Milestone) (models.Milestone, error)
	GetMilestone(id string) (models.Milestone, error)
	ListMilestone() ([]models.Milestone, error)
}

// MilestoneMemory is a thread-safe in-memory MilestoneStore.
type MilestoneMemory struct {
	mu    sync.RWMutex
	items map[string]models.Milestone
}

// NewMilestoneMemory returns an initialised store.
func NewMilestoneMemory() *MilestoneMemory {
	return &MilestoneMemory{items: make(map[string]models.Milestone)}
}

func (m *MilestoneMemory) CreateMilestone(v models.Milestone) (models.Milestone, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *MilestoneMemory) GetMilestone(id string) (models.Milestone, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Milestone{}, ErrNotFound
	}
	return v, nil
}

func (m *MilestoneMemory) ListMilestone() ([]models.Milestone, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Milestone, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}