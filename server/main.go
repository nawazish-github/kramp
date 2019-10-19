package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nawazish-github/kramp/handlers"
)

//ToDo: Config file for no of response elements

func main() {
	r := gin.Default()
	r.GET("/query/:searchString", handlers.QueryHandler)
	log.Fatal(r.Run())
}
