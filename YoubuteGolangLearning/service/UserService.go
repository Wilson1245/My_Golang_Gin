package service

import (
	db "golangAPI/database"
	session "golangAPI/middlewares"
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
	isDelete := pojo.DeleteUser(c.Param("id"))
	if isDelete {
		c.JSON(http.StatusOK, "Successfully")
		return
	}
	c.JSON(http.StatusNotFound, "Error")
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

func CreateUsers(c *gin.Context) {
	users := pojo.Users{}
	err := c.BindJSON(&users)
	if err != nil {
		c.String(400, "Error:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// LoginUser
func LoginUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := pojo.LoginUser(name, password)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	session.SaveSession(c, user.Id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"user":    user,
		"Session": session.GetSession(c),
	})
}

// User get Session
func GetSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": session.GetSession(c),
	})
}

// Logout User
func LogoutUser(c *gin.Context) {
	session.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}

// Redis User
func RedisUser(c *gin.Context) {
	id := c.Param("id")
	user := pojo.User{}
	db.DBconnect.Find(&user, id)
	c.Set("dbResult", user)
}

// Redis User All
func RedisUserAll(c *gin.Context) {
	users := []pojo.User{}
	db.DBconnect.Find(&users)
	c.Set("dbUserAll", users)
}
