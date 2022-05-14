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

func GenerateWorkouts() {
	for i := 0; i < 10; i++ {
		workout := DB.Workout{
			Name:     fmt.Sprintf("Test Workout %d", i),
			Calories: i * 100,
			Minutes:  i * 5,
		}
		Config.DB.Create(&workout)
	}
}

func GenerateUserWorkouts(userId int) {
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
