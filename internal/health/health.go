// Package health aggregates readiness checks.
package health

// Check is a named readiness probe.
type Check struct {
	Name string
	Fn   func() error
}

// Report is the aggregated result of running checks.
type Report struct {
	Healthy bool              `json:"healthy"`
	Details map[string]string `json:"details"`
}

// Run executes all checks and returns a report.
func Run(checks ...Check) Report {
	r := Report{Healthy: true, Details: map[string]string{}}
	for _, c := range checks {
		if err := c.Fn(); err != nil {
			r.Healthy = false
			r.Details[c.Name] = err.Error()
		} else {
			r.Details[c.Name] = "ok"
		}
	}
	return r
}