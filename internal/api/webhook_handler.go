package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// WebhookHandler serves webhook endpoints.
type WebhookHandler struct {
	store store.WebhookStore
	seq   int
}

// NewWebhookHandler constructs a handler backed by s.
func NewWebhookHandler(s store.WebhookStore) *WebhookHandler { return &WebhookHandler{store: s} }

type createWebhookRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/webhooks.
func (h *WebhookHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createWebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Webhook{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateWebhook(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create webhook")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/webhooks.
func (h *WebhookHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListWebhook()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list webhooks")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/webhooks/{id}.
func (h *WebhookHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetWebhook(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "webhook not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}