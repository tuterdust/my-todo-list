package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tuterdust/my-todo-list/database"
)

var (
	dbManager *database.DBManager
	logger    *log.Logger
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/list", getAllToDoListHandler)
	r.GET("/list/:listUUID/:taskUUID", getTaskInToDoListHandler)
	r.GET("/list/:listUUID", getToDoListHandler)
	r.POST("/list", createToDoListHandler)
	r.POST("/task", createTaskHandler)
	r.PUT("/list", updateListHandler)
	r.PUT("/task", updateTaskHandler)
	r.DELETE("/list/:listUUID", deleteListHandler)
	r.DELETE("/task/:taskUUID", deleteTaskHandler)

	return r
}

func main() {
	setLogFiles()
	r := setupRouter()
	dbManager = database.NewDBManager()
	dbManager.Connect()
	r.Run(":8080")
}

func setLogFiles() {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", os.ModePerm)
	}
	setGinLog()
	setErrorLog()
}

func setGinLog() {
	gin.DisableConsoleColor()
	f, _ := os.Create("log/gin_info.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func setErrorLog() {
	f, err := os.OpenFile("log/error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger = log.New(f, "API  ", log.LstdFlags)
	logger.Println("Error log starts")
}
