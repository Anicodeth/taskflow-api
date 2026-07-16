// Package page provides simple offset/limit pagination helpers.
package page

// Clamp normalises limit and offset into safe bounds.
func Clamp(limit, offset, max int) (int, int) {
	if limit <= 0 {
		limit = 20
	}
	if limit > max {
		limit = max
	}
	if offset < 0 {
		offset = 0
	}
	return limit, offset
}

// Slice returns the window of items described by limit and offset.
func Slice[T any](items []T, limit, offset int) []T {
	if offset >= len(items) {
		return []T{}
	}
	end := offset + limit
	if end > len(items) {
		end = len(items)
	}
	return items[offset:end]
}