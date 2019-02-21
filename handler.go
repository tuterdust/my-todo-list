package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tuterdust/my-todo-list/model"
)

func getAllToDoListHandler(c *gin.Context) {
	allList := model.NewList()
	c.JSON(http.StatusOK, gin.H{"list": allList})
}

func getAllTaskInToDoListHandler(c *gin.Context) {
	allTask := make([]*model.Task, 0)
	if err := dbManager.FetchAllTask(&allTask); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	listID := c.Params.ByName("listID")
	c.JSON(http.StatusOK, gin.H{
		"list_id": listID,
		"tasks":   allTask})
}

func getTaskInToDoListHandler(c *gin.Context) {
	task := model.NewTask()
	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	if err := dbManager.FetchTaskFromID(taskID, task); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}
