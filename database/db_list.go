package database

import (
	"github.com/tuterdust/my-todo-list/model"
)

// FetchAllList fetchs every TODO list from the database (exclude task within the lists)
func (dbManager *DBManager) FetchAllList(
	lists *[]*model.List,
) error {
	s := `	SELECT
				*
			FROM
				"list";`
	return dbManager.db.Select(lists, s)
}

// FetchListInfoByID fetchs TODO list information(without Task) from the database using list ID
func (dbManager *DBManager) FetchListInfoByID(
	id int,
	list *model.List,
) error {
	s := `	SELECT
				*
			FROM
				"list"
			WHERE "list"."id" = $1;`
	return dbManager.db.Get(list, s, id)
}

// FetchDetailedListByID fetchs TODO list information (with Task) from the database using list ID
func (dbManager *DBManager) FetchDetailedListByID(
	id int,
	list *model.List,
) error {
	if err := dbManager.FetchListInfoByID(id, list); err != nil {
		return err
	}
	tasks := make([]*model.Task, 0)
	if err := dbManager.FetchAllTask(id, &tasks); err != nil {
		return err
	}
	list.Tasks = &tasks
	return nil
}

// CreateNewList creates new TODO list record in the database
func (dbManager *DBManager) CreateNewList(
	listName, owner string,
) error {
	s := `INSERT INTO "public"."list" ("name", "owner")
		  	VALUES ($1, $2);`
	row, err := dbManager.db.Query(s, listName, owner)
	defer row.Close()
	if err != nil {
		return err
	}
	return nil
}
