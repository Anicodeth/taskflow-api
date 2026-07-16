package store

import (
	"sync"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// WebhookStore describes persistence operations for webhook resources.
type WebhookStore interface {
	CreateWebhook(v models.Webhook) (models.Webhook, error)
	GetWebhook(id string) (models.Webhook, error)
	ListWebhook() ([]models.Webhook, error)
}

// WebhookMemory is a thread-safe in-memory WebhookStore.
type WebhookMemory struct {
	mu    sync.RWMutex
	items map[string]models.Webhook
}

// NewWebhookMemory returns an initialised store.
func NewWebhookMemory() *WebhookMemory {
	return &WebhookMemory{items: make(map[string]models.Webhook)}
}

func (m *WebhookMemory) CreateWebhook(v models.Webhook) (models.Webhook, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[v.ID] = v
	return v, nil
}

func (m *WebhookMemory) GetWebhook(id string) (models.Webhook, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.items[id]
	if !ok {
		return models.Webhook{}, ErrNotFound
	}
	return v, nil
}

func (m *WebhookMemory) ListWebhook() ([]models.Webhook, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]models.Webhook, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}