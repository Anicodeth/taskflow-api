package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// ActivityHandler serves activity endpoints.
type ActivityHandler struct {
	store store.ActivityStore
	seq   int
}

// NewActivityHandler constructs a handler backed by s.
func NewActivityHandler(s store.ActivityStore) *ActivityHandler { return &ActivityHandler{store: s} }

type createActivityRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/activities.
func (h *ActivityHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createActivityRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Activity{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateActivity(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create activity")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/activities.
func (h *ActivityHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListActivity()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list activities")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/activities/{id}.
func (h *ActivityHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetActivity(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "activity not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}