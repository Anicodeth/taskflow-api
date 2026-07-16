// Package validate offers minimal field validation helpers.
package validate

import (
	"fmt"
	"strings"
)

// Required returns an error if value is empty after trimming.
func Required(field, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", field)
	}
	return nil
}

// MaxLen returns an error if value exceeds n runes.
func MaxLen(field, value string, n int) error {
	if len([]rune(value)) > n {
		return fmt.Errorf("%s must be at most %d characters", field, n)
	}
	return nil
}