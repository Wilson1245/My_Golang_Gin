package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBconnect *gorm.DB

var err error

func DD() {
	dsn := "root:00000000@tcp(127.0.0.1:3306)/Golang_GinTest?charset=utf8mb4&parseTime=True&loc=Local"
	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
