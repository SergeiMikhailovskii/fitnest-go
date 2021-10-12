package main

import (
	"TestProject/Config"
	"TestProject/Models"
	"TestProject/Routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	url := Config.DbURL(Config.BuildDBConfig())
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	Config.DB = db
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	err = r.Run(":8080")
}
