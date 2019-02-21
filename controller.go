package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllToDoListHandler(c *gin.Context) {
	var allList []*string
	c.JSON(http.StatusOK, gin.H{"list": allList})
}

func getAllTaskInToDoListHandler(c *gin.Context) {
	listID := c.Params.ByName("listID")
	c.JSON(http.StatusOK, gin.H{"list_id": listID})
}
