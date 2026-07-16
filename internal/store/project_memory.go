package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// ProjectStore describes persistence operations for project resources.
type ProjectStore interface {
	CreateProject(v models.Project) (models.Project, error)
	GetProject(id string) (models.Project, error)
	ListProject() ([]models.Project, error)
}

// ProjectMemory is a thread-safe in-memory ProjectStore.
type ProjectMemory struct {
	mu    sync.RWMutex
	items map[string]models.Project
}

// NewProjectMemory returns an initialised store.
func NewProjectMemory() *ProjectMemory {
	return &ProjectMemory{items: make(map[string]models.Project)}
}

func (m *ProjectMemory) CreateProject(v models.Project) (models.Project, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *ProjectMemory) GetProject(id string) (models.Project, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Project{}, ErrNotFound
	}
	return v, nil
}

func (m *ProjectMemory) ListProject() ([]models.Project, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Project, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}