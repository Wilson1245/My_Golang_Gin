package src

import (
	session "golangAPI/middlewares"
	"golangAPI/pojo"
	"golangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users", session.SetSession())

	user.GET("/", service.CachDecoratorUserAll(service.RedisUserAll, "user_all", pojo.User{}))
	user.GET("/:id", service.CachDecorator(service.RedisUser, "id", "user_%s", pojo.User{}))
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUsers)

	// put user
	user.PUT("/:id", service.PutUser)

	//LoginUser
	user.POST("/login", service.LoginUser)

	user.Use(session.AuthSession())
	{
		// delete user
		user.DELETE("/:id", service.DeleteUser)
		// Get user session
		user.GET("/session", service.GetSession)
		// Logout user
		user.GET("/logout", service.LogoutUser)
	}
}
