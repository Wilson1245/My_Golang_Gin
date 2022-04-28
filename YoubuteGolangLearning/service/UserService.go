package service

import (
	"golangAPI/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get User
func GetAllUser(c *gin.Context) {
	users := pojo.FindAllUsers()
	log.Println("Users -> ", users)
	c.JSON(http.StatusOK, users)
}

// Get One User
func GetOneUser(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User -> ", user)
	c.JSON(http.StatusOK, user)
}

// Post User

func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, "錯誤訊息:%s", err.Error())
		return
	}
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusCreated, newUser)
}

// delete User

func DeleteUser(c *gin.Context) {
	user := pojo.DeleteUser(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
}

// put User

func PutUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	user = pojo.UpdateUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}
