package model

import "time"

// Task is a struct representing a task of TODO list
type Task struct {
	ID          int       `db:"id" json:"task_id"`
	Name        string    `db:"name" json:"task_name"`
	Description string    `db:"description" json:"description"`
	Done        bool      `db:"done" json:"done"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// NewTask returns empty Task structs
func NewTask() *Task {
	return &Task{}
}
