package main

import (
	database "golangAPI/database"
	"golangAPI/middlewares"
	src "golangAPI/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLoggerOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLoggerOutput() // setup logging

	router := gin.Default()
	router.Use(gin.Recovery(), middlewares.Logger()) // logging

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	go func() {
		database.DD()
	}()

	router.Run(":8000")
}
