package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	// import postgres driver
	_ "github.com/lib/pq"
)

// DBManager is a struct representing database manager
type DBManager struct {
	db *sqlx.DB
}

// NewDBManager returns new empty instance of DBManager
func NewDBManager() *DBManager {
	return &DBManager{}
}

// Connect connects to database and store in db variable
func (dbManager *DBManager) Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"))

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	dbManager.db = db
}
