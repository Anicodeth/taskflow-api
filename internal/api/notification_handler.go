package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// NotificationHandler serves notification endpoints.
type NotificationHandler struct {
	store store.NotificationStore
	seq   int
}

// NewNotificationHandler constructs a handler backed by s.
func NewNotificationHandler(s store.NotificationStore) *NotificationHandler { return &NotificationHandler{store: s} }

type createNotificationRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/notifications.
func (h *NotificationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createNotificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Notification{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateNotification(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create notification")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/notifications.
func (h *NotificationHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListNotification()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list notifications")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/notifications/{id}.
func (h *NotificationHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetNotification(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "notification not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}