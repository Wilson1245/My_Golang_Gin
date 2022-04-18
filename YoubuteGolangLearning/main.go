package main

import (
	database "golangAPI/database"
	"golangAPI/middlewares"
	src "golangAPI/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(gin.Recovery(), middlewares.Logger())

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	go func() {
		database.DD()
	}()

	router.Run(":8000")
}
