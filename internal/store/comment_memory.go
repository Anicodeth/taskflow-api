package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// CommentStore describes persistence operations for comment resources.
type CommentStore interface {
	CreateComment(v models.Comment) (models.Comment, error)
	GetComment(id string) (models.Comment, error)
	ListComment() ([]models.Comment, error)
}

// CommentMemory is a thread-safe in-memory CommentStore.
type CommentMemory struct {
	mu    sync.RWMutex
	items map[string]models.Comment
}

// NewCommentMemory returns an initialised store.
func NewCommentMemory() *CommentMemory {
	return &CommentMemory{items: make(map[string]models.Comment)}
}

func (m *CommentMemory) CreateComment(v models.Comment) (models.Comment, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *CommentMemory) GetComment(id string) (models.Comment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Comment{}, ErrNotFound
	}
	return v, nil
}

func (m *CommentMemory) ListComment() ([]models.Comment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Comment, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}