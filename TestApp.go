package main

import (
	"TestProject/Config"
	"TestProject/Models"
	"TestProject/Routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	Config.DB, _ = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("Status:", err)
	//}
	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
