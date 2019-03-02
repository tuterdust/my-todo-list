package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// Task is a struct representing a task of TODO list
type Task struct {
	UUID        uuid.UUID `db:"uuid" json:"task_uuid"`
	ListUUID    uuid.UUID `db:"list_uuid" json:"-"`
	Name        string    `db:"name" json:"task_name"`
	Description string    `db:"description" json:"description"`
	Done        bool      `db:"status" json:"done"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// NewTask returns empty Task structs
func NewTask() *Task {
	return &Task{}
}
