// Package slug converts arbitrary text into url-safe slugs.
package slug

import (
	"strings"
	"unicode"
)

// Make lower-cases s and replaces non-alphanumeric runs with single hyphens.
func Make(s string) string {
	var b strings.Builder
	prevDash := false
	for _, r := range strings.ToLower(strings.TrimSpace(s)) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(r)
			prevDash = false
		} else if !prevDash {
			b.WriteRune('-')
			prevDash = true
		}
	}
	return strings.Trim(b.String(), "-")
}