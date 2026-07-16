package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// AttachmentHandler serves attachment endpoints.
type AttachmentHandler struct {
	store store.AttachmentStore
	seq   int
}

// NewAttachmentHandler constructs a handler backed by s.
func NewAttachmentHandler(s store.AttachmentStore) *AttachmentHandler { return &AttachmentHandler{store: s} }

type createAttachmentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/attachments.
func (h *AttachmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createAttachmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Attachment{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateAttachment(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create attachment")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/attachments.
func (h *AttachmentHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListAttachment()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list attachments")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/attachments/{id}.
func (h *AttachmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetAttachment(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "attachment not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}