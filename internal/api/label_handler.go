package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// LabelHandler serves label endpoints.
type LabelHandler struct {
	store store.LabelStore
	seq   int
}

// NewLabelHandler constructs a handler backed by s.
func NewLabelHandler(s store.LabelStore) *LabelHandler { return &LabelHandler{store: s} }

type createLabelRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/labels.
func (h *LabelHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createLabelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Label{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateLabel(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create label")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/labels.
func (h *LabelHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListLabel()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list labels")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/labels/{id}.
func (h *LabelHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetLabel(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "label not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}