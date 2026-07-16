// Package httpx contains small HTTP helper utilities.
package httpx

import (
	"encoding/json"
	"net/http"
)

// DecodeJSON reads and decodes a JSON request body into v.
func DecodeJSON(r *http.Request, v any) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}