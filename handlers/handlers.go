package handlers

import (
	"log"

	"github.com/nawazish-github/kramp/models"

	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/kramp/cache"
	"github.com/nawazish-github/kramp/executor"
)

func QueryHandler(ginContext *gin.Context) {
	searchString := ginContext.Param("searchString")
	log.Println("Search String: ", searchString)
	if searchString == "" {
		ginContext.JSON(200, gin.H{
			"message": "Empty search string",
		})
	}

	var exec executor.IExecutor
	exec = &executor.SimpleExecutor{}
	response := exec.Fetch(searchString)

	handleCachingIfRequired(&response)
	boundResponseIfRequired(&response)

	ginContext.JSON(200, response)
}

func handleCachingIfRequired(response *models.Response) {
	c := cache.GetCacheInstance()
	if len(response.Albums.Results) > 0 && len(response.Books.Items) > 0 {
		c.UpdateCache(response)
	}
	if &response.Albums == nil || len(response.Albums.Results) == 0 {
		response.Albums = c.Albums
	}
	if &response.Books == nil || len(response.Books.Items) == 0 {
		response.Books = c.Books
	}
}

func boundResponseIfRequired(response *models.Response) {
	config := models.GetConfigInstance()
	if len(response.Books.Items) > config.MaxBooks {
		response.Books.Items = response.Books.Items[0:config.MaxBooks]
	}
	if len(response.Albums.Results) > config.MaxAlbums {
		response.Albums.Results = response.Albums.Results[0:config.MaxAlbums]
	}
}
