// Package ptr provides helpers for working with pointers to values.
package ptr

// Of returns a pointer to v.
func Of[T any](v T) *T { return &v }

// Or dereferences p or returns fallback when p is nil.
func Or[T any](p *T, fallback T) T {
	if p == nil {
		return fallback
	}
	return *p
}