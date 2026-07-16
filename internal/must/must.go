// Package must provides panic-on-error helpers for initialisation code.
package must

// Get returns v or panics if err is non-nil.
func Get[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}