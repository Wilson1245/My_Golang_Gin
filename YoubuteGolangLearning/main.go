package main

import (
	database "golangAPI/database"
	src "golangAPI/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	go func() {
		database.DD()
	}()

	router.Run(":8000")
}
