package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// UserStore describes persistence operations for user resources.
type UserStore interface {
	CreateUser(v models.User) (models.User, error)
	GetUser(id string) (models.User, error)
	ListUser() ([]models.User, error)
}

// UserMemory is a thread-safe in-memory UserStore.
type UserMemory struct {
	mu    sync.RWMutex
	items map[string]models.User
}

// NewUserMemory returns an initialised store.
func NewUserMemory() *UserMemory {
	return &UserMemory{items: make(map[string]models.User)}
}

func (m *UserMemory) CreateUser(v models.User) (models.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *UserMemory) GetUser(id string) (models.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.User{}, ErrNotFound
	}
	return v, nil
}

func (m *UserMemory) ListUser() ([]models.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.User, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}