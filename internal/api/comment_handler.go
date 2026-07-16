package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// CommentHandler serves comment endpoints.
type CommentHandler struct {
	store store.CommentStore
	seq   int
}

// NewCommentHandler constructs a handler backed by s.
func NewCommentHandler(s store.CommentStore) *CommentHandler { return &CommentHandler{store: s} }

type createCommentRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/comments.
func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.Comment{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateComment(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create comment")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/comments.
func (h *CommentHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListComment()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list comments")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/comments/{id}.
func (h *CommentHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetComment(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "comment not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}