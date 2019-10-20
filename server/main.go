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
	r := gin.Default()
	r.GET("/query/:searchString", handlers.QueryHandler)
	log.Fatal(r.Run())
}
