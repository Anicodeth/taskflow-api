// Package clock provides an injectable time source for testable code.
package clock

import "time"

// Clock returns the current time; swap it in tests for determinism.
type Clock func() time.Time

// System is the default wall-clock implementation.
func System() time.Time { return time.Now().UTC() }