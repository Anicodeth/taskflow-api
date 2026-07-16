// Package set implements a minimal generic set type.
package set

// Set is an unordered collection of unique comparable values.
type Set[T comparable] map[T]struct{}

// New builds a Set from items.
func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, it := range items {
		s[it] = struct{}{}
	}
	return s
}

// Add inserts v.
func (s Set[T]) Add(v T) { s[v] = struct{}{} }

// Has reports membership.
func (s Set[T]) Has(v T) bool { _, ok := s[v]; return ok }

// Len returns the size.
func (s Set[T]) Len() int { return len(s) }