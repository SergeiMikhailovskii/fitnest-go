package Dashboard

import (
	"TestProject/Config"
	"TestProject/Models/PrivateArea/DB"
	"fmt"
	"time"
)

func GenerateNotificationsStub(userId int) {
	for i := 0; i < 10; i++ {
		notification := DB.Notification{
			UserId: userId,
			Text:   fmt.Sprintf("Test Notification %d", i),
			Date:   time.Now(),
		}
		Config.DB.Create(&notification)
	}
}

func GenerateWorkoutsStub() {
	for i := 0; i < 10; i++ {
		workout := DB.Workout{
			Name:     fmt.Sprintf("Test Workout %d", i),
			Calories: i * 100,
			Minutes:  i * 5,
		}
		Config.DB.Create(&workout)
	}
}

func GenerateUserWorkoutsStub(userId int) {
	for i := 1; i <= 5; i++ {
		userWorkout := DB.UserWorkout{
			UserId:    userId,
			WorkoutId: i,
			Progress:  0.1 * float32(i),
		}
		err := Config.DB.Create(&userWorkout).Error
		if err != nil {
			panic(err)
		}
	}
}

func GenerateWaterIntakeStub(userId int) {
	for i := 1; i <= 20; i++ {
		waterIntake := DB.WaterIntake{
			UserId: userId,
			Time:   time.Now().Add(-time.Hour * time.Duration(i)),
			Amount: 100 * i,
		}
		err := Config.DB.Create(&waterIntake).Error
		if err != nil {
			panic(err)
		}
	}
}

func GenerateCaloriesIntakeStub(userId int) {
	for i := 1; i <= 20; i++ {
		caloriesIntake := DB.CaloriesIntake{
			UserId: userId,
			Time:   time.Now().Add(-time.Hour * time.Duration(i)),
			Amount: i,
		}
		err := Config.DB.Create(&caloriesIntake).Error
		if err != nil {
			panic(err)
		}
	}
}

func GenerateWaterIntakeAimStub(userId int) {
	waterIntakeAim := DB.ActivityAim{
		UserId:            userId,
		WaterIntakeAmount: 4000,
		CaloriesAmount:    1000,
	}
	err := Config.DB.Create(&waterIntakeAim).Error
	if err != nil {
		panic(err)
	}
}

func GenerateSleepTimeStub(userId int) {
	sleepTime := DB.SleepTime{
		UserId: userId,
		From:   time.Now().Add(-time.Hour * 8),
		To:     time.Now(),
	}
	err := Config.DB.Create(&sleepTime).Error
	if err != nil {
		panic(err)
	}
}

func GenerateHeartRateStub(userId int) {
	heartRate := DB.HeartRate{
		UserId: userId,
		Rate:   78,
		Date:   time.Now(),
	}
	err := Config.DB.Create(&heartRate).Error
	if err != nil {
		panic(err)
	}
}
