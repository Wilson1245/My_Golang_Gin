package main

import (
	. "golangAPI/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	AddUserRouter(v1)
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message":  "ping",
	// 		"message2": "Success!",
	// 	})
	// })
	// router.POST("/ping/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.JSON(200, gin.H{
	// 		"id": id,
	// 	})
	// })
	router.Run(":8000")
}
