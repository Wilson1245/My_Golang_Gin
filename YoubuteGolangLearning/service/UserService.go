package service

import (
	"golangAPI/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var userList = []pojo.User{}

// Get User
func GetAllUser(c *gin.Context) {
	// c.JSON(http.StatusOK, userList)
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
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	// userList = append(userList, user)
	// c.JSON(http.StatusOK, "Successfully posted")
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusCreated, newUser)
}

// delete User

func DeleteUser(c *gin.Context) {
	// userId, _ := strconv.Atoi(c.Param("id"))
	// for _, user := range userList {
	// 	log.Println(user)

	// 	userList = append(userList[:userId], userList[userId+1:]...)
	// 	c.JSON(http.StatusOK, "Successfully deleted")
	// 	return

	// }
	// c.JSON(http.StatusNotFound, "Error")
	user := pojo.DeleteUser(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
}

// put User

func PutUser(c *gin.Context) {
	// beforeUser := pojo.User{}
	// err := c.BindJSON(&beforeUser)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "Error")
	// }
	// userId, _ := strconv.Atoi(c.Param("id"))
	// for key, user := range userList {
	// 	if userId == user.Id {
	// 		userList[key] = beforeUser
	// 		log.Println(userList[key])
	// 		c.JSON(http.StatusOK, "Successfully")
	// 		return
	// 	}
	// }
	// c.JSON(http.StatusNotFound, "Error")
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
