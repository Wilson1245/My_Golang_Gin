package pojo

import (
	db "golangAPI/database"
)

type User struct {
	Id       int    `json:"UserId", binding:"omitempty"`
	Name     string `json:"UserName", binding:"gt=5"`
	Password string `json:"UserPassword", binding:"min=4, max=20"`
	Email    string `json:"UserEmail", binding:"required"`
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
func DeleteUser(userId string) User {
	user := User{}
	db.DBconnect.Where("id = ?", userId).Delete(&user)
	return user
}

// UpdateUser
func UpdateUser(userId string, user User) User {
	db.DBconnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}
