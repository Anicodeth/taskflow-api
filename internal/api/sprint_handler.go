package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// SprintHandler serves sprint endpoints.
type SprintHandler struct {
	store store.SprintStore
	seq   int
}

// NewSprintHandler constructs a handler backed by s.
func NewSprintHandler(s store.SprintStore) *SprintHandler { return &SprintHandler{store: s} }

type createSprintRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/sprints.
func (h *SprintHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createSprintRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Sprint{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateSprint(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create sprint")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/sprints.
func (h *SprintHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListSprint()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list sprints")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/sprints/{id}.
func (h *SprintHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetSprint(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "sprint not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}