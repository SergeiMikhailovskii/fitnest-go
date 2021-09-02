package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DB_USERNAME = "Sergei"
	DB_PASSWORD = "12345"
	DB_NAME     = "go_test"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = "3306"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return nil
	}

	return db
}
