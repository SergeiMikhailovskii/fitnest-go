package main

import (
	"TestProject/Config"
	"TestProject/Models"
	"TestProject/Models/Onboarding"
	"TestProject/Models/PrivateArea/DB"
	"TestProject/Models/Registration"
	"TestProject/Routes"
	"TestProject/Util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var PORT = "8080"

func main() {
	initializeDB()
	r := Routes.SetupRouter()
	port := os.Getenv("PORT")
	if Util.IsEmpty(port) {
		port = PORT
	}
	_ = r.Run(":" + port)
}

func initializeDB() {
	url := Config.DbURL(Config.BuildDBConfig())
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	Config.DB = db
	_ = Config.DB.AutoMigrate(
		&Models.User{},
		&Onboarding.Onboarding{},
		&Registration.PrimaryInfo{},
		&Registration.AnthropometryModel{},
		&Registration.GoalModel{},
		&DB.Notification{},
		&DB.HeartRate{},
		&DB.Workout{},
		&DB.UserWorkout{},
		&DB.WaterIntake{},
	)
}
