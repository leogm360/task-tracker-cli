package tasks

import (
	"encoding/json"
	"fmt"
	"time"
)

// Represents the current task status. It can be either of todo (0),
// in-progress (1) or done (2)
type TaskStatus int

const (
	// Represent the state of Task entity the its yet to be worked upon
	Todo TaskStatus = iota
	// Represent the state of Task entity thats its beign worked upon
	InProgress
	// Repesent the state of Task entity thats has been concluded
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
		panic(fmt.Errorf("invalid task status: %d", ts))
	}
}

// Represents a Task entity containing the id, descrition, status, creation time
// and last update time
type Task struct {
	Id          uint       `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

// Creates a new task entity with the default values, validating the input
func NewTask(id uint, description string) *Task {
	return &Task{
		Id:          id,
		Description: description,
		Status:      0,
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
}

// Serialize Task entity into JSON, converting TaskStatus from int to its string
// representation
func (t *Task) MarshalJSON() ([]byte, error) {
	type TaskMarshal Task

	return json.Marshal(&struct {
		*TaskMarshal
		Status string `json:"status"`
	}{
		TaskMarshal: (*TaskMarshal)(t),
		Status:      t.Status.String(),
	})
}
