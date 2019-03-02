package model

import (
	"time"

	"github.com/satori/go.uuid"
)

// List is a struct representing a TODO list
type List struct {
	UUID      uuid.UUID `db:"uuid" json:"list_uuid"`
	Name      string    `db:"name" json:"list_name"`
	Owner     string    `db:"owner" json:"owner"`
	Tasks     *[]*Task  `json:"tasks"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// NewList returns empty List structs
func NewList() *List {
	return &List{}
}
