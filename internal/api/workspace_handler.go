package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// WorkspaceHandler serves workspace endpoints.
type WorkspaceHandler struct {
	store store.WorkspaceStore
	seq   int
}

// NewWorkspaceHandler constructs a handler backed by s.
func NewWorkspaceHandler(s store.WorkspaceStore) *WorkspaceHandler { return &WorkspaceHandler{store: s} }

type createWorkspaceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/workspaces.
func (h *WorkspaceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createWorkspaceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Workspace{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateWorkspace(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create workspace")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/workspaces.
func (h *WorkspaceHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListWorkspace()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list workspaces")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/workspaces/{id}.
func (h *WorkspaceHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetWorkspace(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "workspace not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}