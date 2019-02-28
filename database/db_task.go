package database

import "github.com/tuterdust/my-todo-list/model"

// FetchAllTask fetchs every task in the list from the database
func (dbManager *DBManager) FetchAllTask(
	listID int,
	tasks *[]*model.Task) error {
	s := `	SELECT
				*
			FROM
				"task"
			WHERE
				"task"."list_id" = $1;`
	return dbManager.db.Select(tasks, s, listID)
}

//FetchTaskFromID fetchs a task from the database using it's ID (and list ID)
func (dbManager *DBManager) FetchTaskFromID(
	id, listID int,
	task *model.Task) error {
	s := `	SELECT
				*
			FROM
				"task"
			WHERE
				"task"."id" = $1
				AND "task"."list_id" = $2;`
	return dbManager.db.Get(task, s, id, listID)
}

// CreateNewTask creates new task record in the database
func (dbManager *DBManager) CreateNewTask(
	listID int,
	name, description string) error {
	s := `	INSERT INTO "public"."task" ("list_id", "name" , "description")
				VALUES ($1, $2 , $3);`
	row, err := dbManager.db.Query(s, listID, name, description)
	defer row.Close()
	if err != nil {
		return err
	}
	return nil
}
