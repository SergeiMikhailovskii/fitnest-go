package main

import (
	"TestProject/Config"
	"TestProject/Models"
	"TestProject/Models/Onboarding"
	"TestProject/Models/PrivateArea/DB"
	"TestProject/Models/Registration"
	"TestProject/Routes"
	"TestProject/Util"
	"errors"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var PORT = "8080"

func main() {
	initializeDB()
	initEmailConfig()
	r := Routes.SetupRouter()
	port := os.Getenv("PORT")
	if Util.IsEmpty(port) {
		port = PORT
	}
	_ = r.Run(":" + port)
}

func initializeDB() {
	url := Config.DbURL(Config.BuildDBConfig())
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		}),
	})
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
		&DB.Steps{},
		&DB.ActivityAim{},
		&DB.SleepTime{},
		&DB.CaloriesIntake{},
	)
}

func initEmailConfig() {
	if _, err := os.Stat("email_config.yaml"); errors.Is(err, os.ErrNotExist) {
		data, errMarshal := yaml.Marshal(Models.EmailConfig{
			Email:    os.Getenv("CONFIG_EMAIL"),
			Password: os.Getenv("CONFIG_PASSWORD"),
			SmtpHost: os.Getenv("CONFIG_SMTP_HOST"),
			SmtpPort: os.Getenv("CONFIG_SMTP_PORT"),
		})
		if errMarshal != nil {
			panic(errMarshal)
		}

		errWrite := ioutil.WriteFile("email_config.yaml", data, os.ModePerm)

		if errWrite != nil {
			panic(errWrite)
		}
	}
}
