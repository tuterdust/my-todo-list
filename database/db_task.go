package database

import (
	uuid "github.com/satori/go.uuid"
	"github.com/tuterdust/my-todo-list/model"
)

// FetchAllTask fetchs every task in the list from the database
func (dbManager *DBManager) FetchAllTask(
	listUUID uuid.UUID,
	tasks *[]*model.Task) error {
	s := `	SELECT
				*
			FROM
				"task"
			WHERE
				"task"."list_uuid" = $1;`
	return dbManager.db.Select(tasks, s, listUUID)
}

//FetchTaskFromID fetchs a task from the database using it's ID (and list ID)
func (dbManager *DBManager) FetchTaskFromID(
	uuid, listUUID uuid.UUID,
	task *model.Task) error {
	s := `	SELECT
				*
			FROM
				"task"
			WHERE
				"task"."uuid" = $1
				AND "task"."list_uuid" = $2;`
	return dbManager.db.Get(task, s, uuid, listUUID)
}

// CreateNewTask creates new task record in the database
func (dbManager *DBManager) CreateNewTask(
	listUUID uuid.UUID,
	name, description string) error {
	s := `	INSERT INTO "public"."task" ("list_uuid", "name" , "description")
				VALUES ($1, $2 , $3);`
	row, err := dbManager.db.Query(s, listUUID, name, description)
	if err != nil {
		return err
	}
	defer row.Close()
	return nil
}

// UpdateTask updates the task record in the database
func (dbManager *DBManager) UpdateTask(
	newName, newDescription string,
	status bool,
	listUUID, taskUUID uuid.UUID) error {
	s := `	UPDATE
				"public"."task"
			SET
				"name" = $1,
				"description" = $2,
				"done" = $3,
				"list_uuid" = $4,
				"updated_at" = NOW()
			WHERE
				"uuid" = $5;`
	row, err := dbManager.db.Query(s, newName, newDescription, status, listUUID, taskUUID)
	if err != nil {
		return err
	}
	defer row.Close()
	return nil
}

// DeleteTask deletes a task from the database using task ID
func (dbManager *DBManager) DeleteTask(
	taskID int,
) error {
	s := ``
	row, err := dbManager.db.Query(s, taskID)
	if err != nil {
		return err
	}
	defer row.Close()
	return nil
}
