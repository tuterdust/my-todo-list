package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tuterdust/my-todo-list/database"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/list", getAllToDoListHandler)
	r.GET("/list/:listID", getAllTaskInToDoListHandler)

	return r
}

func main() {
	r := setupRouter()
	db := database.NewDBManager()
	db.Connect()
	r.Run(":8080")
}
