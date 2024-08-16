package tasks

import (
	"errors"
	"fmt"
	"task-tracker-cli/src/data"
	"time"
)

// Represents the current task status. It can be either of todo (0),
// in-progress(1) or done (2)
type TaskStatus int

const (
	Todo TaskStatus = iota
	InProgress
	Done
)

// Returns the string representation of a task status
func (ts TaskStatus) String() string {
	switch {
	case ts == 0:
		return "todo"
	case ts == 1:
		return "in-progress"
	case ts == 2:
		return "done"
	default:
		panic(fmt.Sprintf("Invalid task status: %d", ts))
	}
}

// Represents a Task entity containing the id, descrition, status, creation time
// and last update time
type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   string     `json:"createdAt"`
	UpdatedAt   string     `json:"updatedAt"`
}

// Creates a new task entity with the default values, validating the input
func NewTask(description string) (*Task, error) {
	return &Task{
		Id:          1,
		Description: description,
		Status:      0,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}, nil
}

// Update the task entity 'description' field, validating the input
func (t *Task) UpdateDescription(newDescription string) (bool, error) {
	if newDescription == "" {
		return false, errors.New("description cannot be empty")
	}

	if len(newDescription) < 4 {
		return false, errors.New("description should be at least 4 characters long")
	}

	t.Description = newDescription
	t.UpdatedAt = time.Now().Format(time.RFC3339)

	return true, nil
}

// Update the task entity 'status' field, validating the input
func (t *Task) UpdateStatus(newStatus TaskStatus) (bool, error) {
	if t.Status == newStatus {
		return false, errors.New("task is already in this status")
	}

	t.Status = newStatus
	t.UpdatedAt = time.Now().Format(time.RFC3339)

	return true, nil
}

// Persists the task entity on the given data source
func (t *Task) Persist(ds data.DataSourcer) (bool, error) {
	ok, err := ds.Write(t)

	return ok, err
}
