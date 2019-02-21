package database

import "github.com/tuterdust/my-todo-list/model"

//FetchTaskFromID fetchs a task from the database using it's ID
func (dbManager *DBManager) FetchTaskFromID(
	id int,
	task *model.Task) error {
	s := `	SELECT
				*
			FROM
				"task"
			WHERE
				"task"."id" = $1;`
	return dbManager.db.Get(task, s, id)
}
