// Package version carries build metadata for the service.
package version

// Version is the current semantic version of the API.
const Version = "0.1.0"

// String returns a human-readable version banner.
func String() string { return "taskflow-api " + Version }