package database

import (
	"github.com/satori/go.uuid"
	"github.com/tuterdust/my-todo-list/src/model"
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
	listUUID uuid.UUID,
	list *model.List,
) error {
	s := `	SELECT
				*
			FROM
				"list"
			WHERE "list"."uuid" = $1;`
	return dbManager.db.Get(list, s, listUUID)
}

// FetchDetailedListByID fetchs TODO list information (with Task) from the database using list ID
func (dbManager *DBManager) FetchDetailedListByID(
	listUUID uuid.UUID,
	list *model.List,
) error {
	if err := dbManager.FetchListInfoByID(listUUID, list); err != nil {
		return err
	}
	tasks := make([]*model.Task, 0)
	if err := dbManager.FetchAllTask(listUUID, &tasks); err != nil {
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
	if err != nil {
		return err
	}
	defer row.Close()
	return nil
}

// UpdateList updates the list record in the database
func (dbManager *DBManager) UpdateList(
	newName, newOwner string,
	listUUID uuid.UUID,
) error {
	s := `	UPDATE
				"public"."list"
			SET
				"name" = $1,
				"owner" = $2,
				"updated_at" = NOW()
			WHERE
				"uuid" = $3;`
	row, err := dbManager.db.Query(s, newName, newOwner, listUUID)
	if err != nil {
		return err
	}
	defer row.Close()
	return nil
}

// DeleteList deletes a list from the database using list ID
func (dbManager *DBManager) DeleteList(
	listUUID uuid.UUID,
) error {
	// delete list
	s := `	DELETE FROM "public"."list"
				WHERE "uuid" = $1;`
	row, err := dbManager.db.Query(s, listUUID)
	if err != nil {
		return err
	}
	defer row.Close()
	return nil
}
