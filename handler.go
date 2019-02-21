package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuterdust/my-todo-list/model"
)

func getAllToDoListHandler(c *gin.Context) {
	allList := model.NewList()
	c.JSON(http.StatusOK, gin.H{"list": allList})
}

func getAllTaskInToDoListHandler(c *gin.Context) {
	allTask := model.NewTask()
	// listID := c.Params.ByName("listID")
	c.JSON(http.StatusOK, gin.H{"list_id": allTask})
}

func getTaskInToDoListHandler(c *gin.Context) {
	task := model.NewTask()
	if err := dbManager.FetchTaskFromID(1, task); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}
