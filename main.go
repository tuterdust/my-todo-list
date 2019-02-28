package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tuterdust/my-todo-list/database"
)

var dbManager *database.DBManager

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/list", getAllToDoListHandler)
	r.GET("/list/:listID/:taskID", getTaskInToDoListHandler)
	r.GET("/list/:listID", getToDoListHandler)
	r.POST("/list", createToDoListHandler)
	r.POST("/task", createTaskHandler)
	r.PUT("/list", updateListHandler)
	r.PUT("/task", updateTaskHandler)

	return r
}

func main() {
	r := setupRouter()
	dbManager = database.NewDBManager()
	dbManager.Connect()
	r.Run(":8080")
}
