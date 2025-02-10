package store

import (
	"sort"
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// Memory is a thread-safe, in-memory TaskStore implementation.
type Memory struct {
	mu    sync.RWMutex
	tasks map[string]models.Task
}

// NewMemory returns an initialised in-memory store.
func NewMemory() *Memory {
	return &Memory{tasks: make(map[string]models.Task)}
}

// CreateTask stores t and returns it.
func (m *Memory) CreateTask(t models.Task) (models.Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tasks[t.ID] = t
	return t, nil
}

// GetTask fetches a task by id.
func (m *Memory) GetTask(id string) (models.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	t, ok := m.tasks[id]
	if !ok {
		return models.Task{}, ErrNotFound
	}
	return t, nil
}

// ListTasks returns all tasks ordered by creation time.
func (m *Memory) ListTasks() ([]models.Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Task, 0, len(m.tasks))
	for _, t := range m.tasks {
		out = append(out, t)
	}
	sort.Slice(out, func(i, j int) bool {
		return out[i].CreatedAt.Before(out[j].CreatedAt)
	})
	return out, nil
}
