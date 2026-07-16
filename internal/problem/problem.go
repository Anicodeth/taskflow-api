// Package problem models typed, HTTP-friendly application errors.
package problem

import "fmt"

// Problem couples a machine code, HTTP status, and human message.
type Problem struct {
	Code    string
	Status  int
	Message string
}

func (p Problem) Error() string { return fmt.Sprintf("%s: %s", p.Code, p.Message) }

// New constructs a Problem.
func New(code string, status int, msg string) Problem {
	return Problem{Code: code, Status: status, Message: msg}
}