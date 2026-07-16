package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// TagHandler serves tag endpoints.
type TagHandler struct {
	store store.TagStore
	seq   int
}

// NewTagHandler constructs a handler backed by s.
func NewTagHandler(s store.TagStore) *TagHandler { return &TagHandler{store: s} }

type createTagRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/tags.
func (h *TagHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createTagRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Tag{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateTag(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create tag")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/tags.
func (h *TagHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListTag()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list tags")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/tags/{id}.
func (h *TagHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetTag(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "tag not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}