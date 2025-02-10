// Package store defines persistence interfaces and implementations.
package store

import (
	"errors"

	"github.com/Anicodeth/taskflow-api/internal/models"
)

// ErrNotFound is returned when a requested record does not exist.
var ErrNotFound = errors.New("record not found")

// TaskStore describes the persistence operations for tasks.
type TaskStore interface {
	CreateTask(t models.Task) (models.Task, error)
	GetTask(id string) (models.Task, error)
	ListTasks() ([]models.Task, error)
}
