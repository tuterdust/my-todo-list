package main

import (
	"net/http"
	"strconv"

	"github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"github.com/tuterdust/my-todo-list/model"
)

func getAllToDoListHandler(c *gin.Context) {
	allList := make([]*model.List, 0)
	if err := dbManager.FetchAllList(&allList); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lists": allList})
}

func getToDoListHandler(c *gin.Context) {
	list := model.NewList()
	listUUID := uuid.FromStringOrNil(c.Params.ByName("listUUID"))

	isDetailed, err := strconv.ParseBool(c.Query("detailed"))
	if err != nil {
		isDetailed = false
	}

	if isDetailed {
		if err := dbManager.FetchDetailedListByID(listUUID, list); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"list": list})
	} else {
		if err := dbManager.FetchListInfoByID(listUUID, list); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"list": list})
	}
}

func getAllTaskInToDoListHandler(c *gin.Context) {
	allTask := make([]*model.Task, 0)
	listUUID := uuid.FromStringOrNil(c.Params.ByName("listUUID"))

	if err := dbManager.FetchAllTask(listUUID, &allTask); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": allTask})
}

func getTaskInToDoListHandler(c *gin.Context) {
	task := model.NewTask()
	taskUUID := uuid.FromStringOrNil(c.Param("taskUUID"))
	listUUID := uuid.FromStringOrNil(c.Params.ByName("listUUID"))

	if err := dbManager.FetchTaskFromID(taskUUID, listUUID, task); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"task": task})
}

func createToDoListHandler(c *gin.Context) {
	listName := c.DefaultPostForm("name", "Untitled List")
	owner := c.PostForm("owner")
	if err := dbManager.CreateNewList(listName, owner); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}

func createTaskHandler(c *gin.Context) {
	listUUID := uuid.FromStringOrNil(c.PostForm("listUUID"))

	taskName := c.DefaultPostForm("name", "Untitled Task")
	description := c.DefaultPostForm("description", "")
	if err := dbManager.CreateNewTask(listUUID, taskName, description); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}

func updateListHandler(c *gin.Context) {
	listUUID := uuid.FromStringOrNil(c.PostForm("listUUID"))

	newName := c.DefaultPostForm("name", "Untitled Task")
	newOwner := c.DefaultPostForm("owner", "")

	if err := dbManager.UpdateList(newName, newOwner, listUUID); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}

func updateTaskHandler(c *gin.Context) {
	listUUID := uuid.FromStringOrNil(c.PostForm("listUUID"))
	taskUUID := uuid.FromStringOrNil(c.PostForm("taskUUID"))

	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	newName := c.DefaultPostForm("name", "Untitled Task")
	newDescription := c.DefaultPostForm("description", "")

	if err := dbManager.UpdateTask(newName, newDescription, status, listUUID, taskUUID); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}

func deleteListHandler(c *gin.Context) {
	listID, err := strconv.Atoi(c.PostForm("listID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := dbManager.DeleteList(listID); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}

func deleteTaskHandler(c *gin.Context) {
	taskID, err := strconv.Atoi(c.PostForm("taskID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := dbManager.DeleteTask(taskID); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}
