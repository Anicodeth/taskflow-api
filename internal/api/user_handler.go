package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// UserHandler serves user endpoints.
type UserHandler struct {
	store store.UserStore
	seq   int
}

// NewUserHandler constructs a handler backed by s.
func NewUserHandler(s store.UserStore) *UserHandler { return &UserHandler{store: s} }

type createUserRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create handles POST /api/users.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}
	h.seq++
	v := models.User{
		ID:          strconv.Itoa(h.seq),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now().UTC(),
	}
	created, err := h.store.CreateUser(v)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create user")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/users.
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	items, err := h.store.ListUser()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list users")
		return
	}
	writeJSON(w, http.StatusOK, items)
}

// Get handles GET /api/users/{id}.
func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	v, err := h.store.GetUser(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}
	writeJSON(w, http.StatusOK, v)
}