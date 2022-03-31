package src

import (
	"golangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", service.GetAllUser)
	user.GET("/:id", service.GetOneUser)
	user.POST("/", service.PostUser)
	// delete user
	user.DELETE("/:id", service.DeleteUser)
	// put user
	user.PUT("/:id", service.PutUser)
}
