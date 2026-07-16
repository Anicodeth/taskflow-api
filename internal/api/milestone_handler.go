package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// MilestoneHandler serves milestone endpoints.
type MilestoneHandler struct {
	store store.MilestoneStore
	seq   int
}

// NewMilestoneHandler constructs a handler backed by s.
func NewMilestoneHandler(s store.MilestoneStore) *MilestoneHandler { return &MilestoneHandler{store: s} }

type createMilestoneRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/milestones.
func (h *MilestoneHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createMilestoneRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Milestone{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateMilestone(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create milestone")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/milestones.
func (h *MilestoneHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListMilestone()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list milestones")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/milestones/{id}.
func (h *MilestoneHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetMilestone(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "milestone not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}