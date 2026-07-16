package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// AttachmentStore describes persistence operations for attachment resources.
type AttachmentStore interface {
	CreateAttachment(v models.Attachment) (models.Attachment, error)
	GetAttachment(id string) (models.Attachment, error)
	ListAttachment() ([]models.Attachment, error)
}

// AttachmentMemory is a thread-safe in-memory AttachmentStore.
type AttachmentMemory struct {
	mu    sync.RWMutex
	items map[string]models.Attachment
}

// NewAttachmentMemory returns an initialised store.
func NewAttachmentMemory() *AttachmentMemory {
	return &AttachmentMemory{items: make(map[string]models.Attachment)}
}

func (m *AttachmentMemory) CreateAttachment(v models.Attachment) (models.Attachment, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *AttachmentMemory) GetAttachment(id string) (models.Attachment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Attachment{}, ErrNotFound
	}
	return v, nil
}

func (m *AttachmentMemory) ListAttachment() ([]models.Attachment, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Attachment, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}