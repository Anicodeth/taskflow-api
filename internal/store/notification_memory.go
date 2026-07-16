package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// NotificationStore describes persistence operations for notification resources.
type NotificationStore interface {
	CreateNotification(v models.Notification) (models.Notification, error)
	GetNotification(id string) (models.Notification, error)
	ListNotification() ([]models.Notification, error)
}

// NotificationMemory is a thread-safe in-memory NotificationStore.
type NotificationMemory struct {
	mu    sync.RWMutex
	items map[string]models.Notification
}

// NewNotificationMemory returns an initialised store.
func NewNotificationMemory() *NotificationMemory {
	return &NotificationMemory{items: make(map[string]models.Notification)}
}

func (m *NotificationMemory) CreateNotification(v models.Notification) (models.Notification, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *NotificationMemory) GetNotification(id string) (models.Notification, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Notification{}, ErrNotFound
	}
	return v, nil
}

func (m *NotificationMemory) ListNotification() ([]models.Notification, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Notification, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}