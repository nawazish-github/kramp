package handlers

import "github.com/gin-gonic/gin"

func QueryHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HELLO GIN",
	})
}
