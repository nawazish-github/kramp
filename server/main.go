package main

import (
	"log"

	"github.com/nawazish-github/kramp/models"

	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/kramp/handlers"
	"github.com/nawazish-github/kramp/json"
)

//ToDo: Config file for no of response elements
var Conf models.Config

func main() {
	Conf = json.UnmarshalConfig()
	router := gin.Default()
	router.GET("/query/:searchString", handlers.QueryHandler)
	router.NoRoute(handlers.NoRouteHandler)
	log.Fatal(router.Run())
}
