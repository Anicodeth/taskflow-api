package api

import (
	"encoding/json"
	"net/http"
)

// errorEnvelope is the standard shape for error responses.
type errorEnvelope struct {
	Error string `json:"error"`
}

// writeJSON serialises v to the response writer with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v != nil {
		_ = json.NewEncoder(w).Encode(v)
	}
}

// writeError writes a JSON error envelope with the given status code.
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, errorEnvelope{Error: msg})
}
