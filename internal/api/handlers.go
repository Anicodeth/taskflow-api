package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Anicodeth/taskflow-api/internal/models"
	"github.com/Anicodeth/taskflow-api/internal/store"
)

// TaskHandler serves task-related endpoints.
type TaskHandler struct {
	store store.TaskStore
	seq   int
}

// NewTaskHandler constructs a TaskHandler backed by s.
func NewTaskHandler(s store.TaskStore) *TaskHandler {
	return &TaskHandler{store: s}
}

type createTaskRequest struct {
	Title  string        `json:"title"`
	Body   string        `json:"body"`
	Status models.Status `json:"status"`
}

// Create handles POST /api/tasks.
func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}
	if req.Title == "" {
		writeError(w, http.StatusUnprocessableEntity, "title is required")
		return
	}
	if req.Status == "" {
		req.Status = models.StatusTodo
	}
	if !req.Status.Valid() {
		writeError(w, http.StatusUnprocessableEntity, "invalid status")
		return
	}
	now := time.Now().UTC()
	h.seq++
	task := models.Task{
		ID:        strconv.Itoa(h.seq),
		Title:     req.Title,
		Body:      req.Body,
		Status:    req.Status,
		CreatedAt: now,
		UpdatedAt: now,
	}
	created, err := h.store.CreateTask(task)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create task")
		return
	}
	writeJSON(w, http.StatusCreated, created)
}

// List handles GET /api/tasks.
func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.store.ListTasks()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not list tasks")
		return
	}
	writeJSON(w, http.StatusOK, tasks)
}

// Get handles GET /api/tasks/{id}.
func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	task, err := h.store.GetTask(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "task not found")
		return
	}
	writeJSON(w, http.StatusOK, task)
}
