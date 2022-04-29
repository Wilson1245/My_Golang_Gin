package main

import (
	database "golangAPI/database"
	"golangAPI/middlewares"
	"golangAPI/pojo"
	src "golangAPI/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLoggerOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLoggerOutput() // setup logging

	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middlewares.UserPasd)
		v.RegisterStructValidation(middlewares.UserList, pojo.Users{})
	}

	router.Use(gin.Recovery(), middlewares.Logger()) // logging

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	go func() {
		database.DD()
	}()

	router.Run(":8000")
}
