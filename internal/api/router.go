package api

import (
	"net/http"

	"github.com/Anicodeth/taskflow-api/internal/store"
)

// NewRouter wires up all routes and middleware and returns the root handler.
func NewRouter(s store.TaskStore) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	tasks := NewTaskHandler(s)
	mux.HandleFunc("GET /api/tasks", tasks.List)
	mux.HandleFunc("POST /api/tasks", tasks.Create)
	mux.HandleFunc("GET /api/tasks/{id}", tasks.Get)

	projectH := NewProjectHandler(store.NewProjectMemory())
	mux.HandleFunc("GET /api/projects", projectH.List)
	mux.HandleFunc("POST /api/projects", projectH.Create)
	mux.HandleFunc("GET /api/projects/{id}", projectH.Get)
	tagH := NewTagHandler(store.NewTagMemory())
	mux.HandleFunc("GET /api/tags", tagH.List)
	mux.HandleFunc("POST /api/tags", tagH.Create)
	mux.HandleFunc("GET /api/tags/{id}", tagH.Get)
	commentH := NewCommentHandler(store.NewCommentMemory())
	mux.HandleFunc("GET /api/comments", commentH.List)
	mux.HandleFunc("POST /api/comments", commentH.Create)
	mux.HandleFunc("GET /api/comments/{id}", commentH.Get)
	// registrations:end

	return Chain(mux, Recovery, Logging)
}
