package pojo

import (
	db "golangAPI/database"
)

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func FindAllUsers() []User {
	users := []User{}
	db.DBconnect.Find(&users)
	return users
}

func FindByUserId(userId string) User {
	user := User{}
	db.DBconnect.Where("id = ?", userId).First(&user)
	return user
}
