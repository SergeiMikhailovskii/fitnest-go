package main

import (
	"TestProject/Config"
	"TestProject/Models"
	"TestProject/Routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	url := Config.DbURL(Config.BuildDBConfig())
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	Config.DB = db
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = r.Run(":" + port)
}
