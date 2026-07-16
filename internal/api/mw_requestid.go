package api

import (
	"context"
	"net/http"
	"strconv"
	"sync/atomic"
)

type ctxKey string

const requestIDKey ctxKey = "request-id"

var reqCounter atomic.Int64

// RequestID assigns a monotonic request id and echoes it via a header.
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := strconv.FormatInt(reqCounter.Add(1), 10)
		w.Header().Set("X-Request-ID", id)
		ctx := context.WithValue(r.Context(), requestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}