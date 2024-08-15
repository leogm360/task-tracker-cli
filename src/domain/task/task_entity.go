package task

import (
	"errors"
	"fmt"
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
	id          int
	description string
	status      TaskStatus
	createdAt   string
	updatedAt   string
}

// Creates a new task entity with the default values, validating the input
func NewTask(description string) (*Task, error) {
	return &Task{
		id:          1,
		description: description,
		status:      0,
		createdAt:   time.Now().Format(time.RFC3339),
		updatedAt:   time.Now().Format(time.RFC3339),
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

	t.description = newDescription
	t.updatedAt = time.Now().Format(time.RFC3339)

	return true, nil
}

// Update the task entity 'status' field, validating the input
func (t *Task) UpdateStatus(newStatus TaskStatus) (bool, error) {
	if t.status == newStatus {
		return false, errors.New("")
	}

	t.status = newStatus
	t.updatedAt = time.Now().Format(time.RFC3339)

	return true, nil
}

// Persists the task entity on the given data source and format
func (t *Task) Persist() (bool, error) {
	panic("Implement persist")
}
