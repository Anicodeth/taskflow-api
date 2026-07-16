package health

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	r := Run(
		Check{Name: "db", Fn: func() error { return nil }},
		Check{Name: "cache", Fn: func() error { return errors.New("down") }},
	)
	if r.Healthy {
		t.Fatal("should be unhealthy")
	}
	if r.Details["db"] != "ok" {
		t.Fatal("db ok")
	}
}