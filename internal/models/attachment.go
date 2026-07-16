package models

import "time"

// Attachment is a TaskFlow attachment resource.
type Attachment struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}