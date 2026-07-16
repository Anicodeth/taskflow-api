// Package retry runs an operation with a fixed number of attempts.
package retry

// Do calls fn up to attempts times, returning the last error.
func Do(attempts int, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
	}
	return err
}