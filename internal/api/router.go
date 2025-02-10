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

	return Chain(mux, Recovery, Logging)
}
