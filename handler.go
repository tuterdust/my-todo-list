package main

import (
	"net/http"
	"strconv"

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
	listID, err := strconv.Atoi(c.Params.ByName("listID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	isDetailed, err := strconv.ParseBool(c.Query("detailed"))
	if err != nil {
		isDetailed = false
	}

	if isDetailed {
		if err := dbManager.FetchDetailedListByID(listID, list); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadGateway, gin.H{"list": list})
	} else {
		if err := dbManager.FetchListInfoByID(listID, list); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"list": list})
	}
}

func getAllTaskInToDoListHandler(c *gin.Context) {
	allTask := make([]*model.Task, 0)
	listID, err := strconv.Atoi(c.Params.ByName("listID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return

	}

	if err := dbManager.FetchAllTask(listID, &allTask); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tasks": allTask})
}

func getTaskInToDoListHandler(c *gin.Context) {
	task := model.NewTask()
	taskID, err := strconv.Atoi(c.Param("taskID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	listID, err := strconv.Atoi(c.Params.ByName("listID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := dbManager.FetchTaskFromID(taskID, listID, task); err != nil {
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
	listID, err := strconv.Atoi(c.PostForm("listID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	taskName := c.DefaultPostForm("name", "Untitled Task")
	description := c.DefaultPostForm("description", "")
	if err := dbManager.CreateNewTask(listID, taskName, description); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}

func updateListHandler(c *gin.Context) {
	listID, err := strconv.Atoi(c.PostForm("listID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	newName := c.DefaultPostForm("name", "Untitled Task")
	newOwner := c.DefaultPostForm("owner", "")

	if err := dbManager.UpdateList(newName, newOwner, listID); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}

func updateTaskHandler(c *gin.Context) {
	listID, err := strconv.Atoi(c.PostForm("listID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	taskID, err := strconv.Atoi(c.PostForm("taskID"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	status, err := strconv.ParseBool(c.PostForm("status"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	newName := c.DefaultPostForm("name", "Untitled Task")
	newDescription := c.DefaultPostForm("description", "")

	list := model.NewList()
	if err := dbManager.FetchListInfoByID(listID, list); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	if err := dbManager.UpdateTask(newName, newDescription, status, listID, taskID); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Success"})
}
