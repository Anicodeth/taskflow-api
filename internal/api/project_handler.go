package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// ProjectHandler serves project endpoints.
type ProjectHandler struct {
	store store.ProjectStore
	seq   int
}

// NewProjectHandler constructs a handler backed by s.
func NewProjectHandler(s store.ProjectStore) *ProjectHandler { return &ProjectHandler{store: s} }

type createProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/projects.
func (h *ProjectHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Project{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateProject(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create project")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/projects.
func (h *ProjectHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListProject()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list projects")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/projects/{id}.
func (h *ProjectHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetProject(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "project not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}