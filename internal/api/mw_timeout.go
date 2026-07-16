package api

import (
	"net/http"
	"time"
)

// Timeout bounds request handling to d using the standard library helper.
func Timeout(d time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, d, `{"error":"request timed out"}`)
	}
}