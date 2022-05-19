package pojo

import (
	db "golangAPI/database"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       int    `json:"UserId" binding:"omitempty"`
	Name     string `json:"UserName" binding:"gt=5"`
	Password string `json:"UserPassword" binding:"min=4,max=20,userpasd"`
	Email    string `json:"UserEmail" binding:"required"`
}

type Users struct {
	UserList     []User `json:"UserList" binding:"gt=0,lt=3"`
	UserListSize int    `json:"UserListSize"`
}

// FindAllUsers
func FindAllUsers() []User {
	users := []User{}
	db.DBconnect.Find(&users)
	return users
}

// FindByUserId
func FindByUserId(userId string) User {
	user := User{}
	db.DBconnect.Where("id = ?", userId).First(&user)
	return user
}

// CreateUser
func CreateUser(user User) User {
	db.DBconnect.Create(&user)
	return user
}

// DeleteUser
func DeleteUser(userId string) bool {
	user := User{}
	result := db.DBconnect.Where("id = ?", userId).Delete(&user)
	return result.RowsAffected > 0
}

// UpdateUser
func UpdateUser(userId string, user User) User {
	db.DBconnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}

// Check if user exists
func LoginUser(name string, password string) User {
	user := User{}
	db.DBconnect.Where("name = ? and password = ?", name, password).First(&user)
	return user
}

// MongoDB --------------------------------

func MgoCreateUser(user User) User {
	db.MgoConnect.Insert(user)
	return user
}

func MgoFindAllUser() []User {
	users := []User{}
	db.MgoConnect.Find(nil).All(&users)
	return users
}

func MgoFindById(id string) User {
	userid, _ := strconv.Atoi(id)
	user := User{}
	db.MgoConnect.Find(bson.M{"id": userid}).One(&user)
	return user
}

func MgoFindByName(name string) User {
	user := User{}
	db.MgoConnect.Find(bson.M{"name": name}).One(&user)
	return user
}
