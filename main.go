package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tuterdust/my-todo-list/database"
)

var dbManager *database.DBManager

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/list", getAllToDoListHandler)
	r.GET("/list/:listID", getAllTaskInToDoListHandler)
	r.GET("/list/:listID/:taskID", getTaskInToDoListHandler)

	return r
}

func main() {
	r := setupRouter()
	dbManager = database.NewDBManager()
	dbManager.Connect()
	r.Run(":8080")
}
