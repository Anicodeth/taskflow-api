// Package idgen produces short, URL-safe unique identifiers.
package idgen

import (
	"crypto/rand"
	"encoding/hex"
)

// New returns a 16-character hex identifier.
func New() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}