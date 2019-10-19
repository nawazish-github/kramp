package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/kramp/executor"
)

func QueryHandler(c *gin.Context) {
	searchString := c.Param("searchString")
	log.Println("Search String: ", searchString)
	if searchString == "" {
		c.JSON(200, gin.H{
			"message": "Empty search string",
		})
	}

	simpleExec := executor.SimpleExecutor{}
	response := simpleExec.Fetch(searchString)
	//ToDo: might have to cache here:
	c.JSON(200, response)
}
