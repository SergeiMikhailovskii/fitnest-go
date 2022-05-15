package Dashboard

import (
	"TestProject/Config"
	"TestProject/Controllers/Registration"
	"TestProject/Models/PrivateArea"
	"TestProject/Models/PrivateArea/DB"
	"TestProject/Models/PrivateArea/Widgets"
	RegistrationModel "TestProject/Models/Registration"
	"github.com/gin-gonic/gin"
	"math"
	"time"
)

func GetDashboardPage(c *gin.Context) (*PrivateArea.Response, error) {
	userId, _ := Registration.GetUserId(c)

	widgetsMap := make(map[string]interface{})

	widgetsMap["HEADER_WIDGET"] = getHeaderWidget(c, userId)
	widgetsMap["BMI_WIDGET"] = getBMIWidget(c)
	widgetsMap["TODAY_TARGET_WIDGET"] = getTodayTargetWidget()
	widgetsMap["ACTIVITY_STATUS_WIDGET"] = getActivityStatusWidget(userId)
	widgetsMap["LATEST_WORKOUT_WIDGET"] = getLatestWorkoutWidget(userId)

	return &PrivateArea.Response{
		Widgets: widgetsMap,
	}, nil
}

func getHeaderWidget(c *gin.Context, userId int) Widgets.HeaderWidget {
	primaryRecord := Registration.GetPrimaryRegistrationRecord(c)
	userName := primaryRecord.FirstName + " " + primaryRecord.LastName

	var activeNotificationsAmount int64
	Config.DB.Model(&DB.Notification{}).
		Where("is_active = ? AND user_id = ?", true, userId).
		Count(&activeNotificationsAmount)

	return Widgets.HeaderWidget{
		Name:          userName,
		Notifications: activeNotificationsAmount,
	}
}

func getBMIWidget(c *gin.Context) *Widgets.BMIWidget {
	userId, _ := Registration.GetUserId(c)

	var anthropometryModel RegistrationModel.AnthropometryModel

	err := Config.DB.Where("user_id = ?", userId).First(&anthropometryModel).Error

	bmiValue := anthropometryModel.Weight / (math.Pow(anthropometryModel.Height/100, 2))
	bmiStatus := getBMIStatusByValue(bmiValue)

	var response *Widgets.BMIWidget
	if err == nil {
		response = &Widgets.BMIWidget{
			Index:  bmiValue,
			Result: bmiStatus,
		}
	} else {
		response = nil
	}

	return response
}

func getTodayTargetWidget() Widgets.TodayTargetWidget {
	return Widgets.TodayTargetWidget{}
}

func getActivityStatusWidget(userId int) Widgets.ActivityStatusWidget {
	heartRate, date, heartRateErr := getLastHeartRate(userId)
	waterIntakeAim, waterIntakeAimErr := getWaterIntakeAim(userId)

	var heartRateWidget *Widgets.HeartRateSubWidget
	var waterIntakeSubWidget *Widgets.WaterIntakeSubWidget

	if heartRateErr == nil {
		heartRateWidget = &Widgets.HeartRateSubWidget{
			Rate: heartRate,
			Date: date,
		}
	} else {
		heartRateWidget = nil
	}

	waterIntakeErr, intakes := getWaterIntakes(userId)

	if waterIntakeAimErr == nil && waterIntakeErr == nil {
		widgetIntakes := mapDBIntakesToWidget(intakes)
		totalTodayWaterIntake := getTotalTodayWaterIntake(intakes)

		waterIntakeSubWidget = &Widgets.WaterIntakeSubWidget{
			Amount:   waterIntakeAim,
			Progress: float32(totalTodayWaterIntake / waterIntakeAim),
			Intakes:  widgetIntakes,
		}
	} else {
		waterIntakeSubWidget = nil
	}

	return Widgets.ActivityStatusWidget{
		HeartRate:   heartRateWidget,
		WaterIntake: waterIntakeSubWidget,
		Sleep:       getSleepDuration(userId),
		Calories: Widgets.CaloriesSubWidget{
			Consumed: 760,
			Left:     230,
		},
	}
}

func getLatestWorkoutWidget(userId int) *Widgets.LatestWorkoutWidget {
	err, lastWorkouts := getLastWorkouts(userId)

	if err != nil {
		return nil
	} else {
		return &Widgets.LatestWorkoutWidget{
			Workouts: lastWorkouts,
		}
	}
}

func getBMIStatusByValue(value float64) string {
	if value < 18.5 {
		return "UNDERWEIGHT"
	} else if value < 25 {
		return "NORMAL_WEIGHT"
	} else if value < 30 {
		return "OVERWEIGHT"
	} else {
		return "OBESITY"
	}
}

func getLastHeartRate(userId int) (int, time.Time, error) {
	var heartRateModel DB.HeartRate
	err := Config.DB.Where("user_id = ?", userId).Last(&heartRateModel).Error
	return heartRateModel.Rate, heartRateModel.Date, err
}

func getWaterIntakeAim(userId int) (int, error) {
	var waterIntakeAim DB.ActivityAim
	err := Config.DB.Where("user_id = ?", userId).Last(&waterIntakeAim).Error
	return waterIntakeAim.WaterIntakeAmount, err
}

func getLastWorkouts(userId int) (error, []Widgets.Workout) {
	rows, err := Config.DB.Model(&DB.UserWorkout{}).Select("user_workout.progress, workout.name, workout.calories, workout.minutes").Joins("inner join workout on workout.id = user_workout.workout_id").Where("user_id = ?", userId).Rows()
	if err != nil {
		return err, nil
	} else {
		var workouts []Widgets.Workout
		for rows.Next() {
			var workout Widgets.Workout
			err = Config.DB.ScanRows(rows, &workout)

			if err != nil {
				return err, nil
			}

			workouts = append(workouts, workout)
		}
		return nil, workouts
	}
}

func getWaterIntakes(userId int) (error, []DB.WaterIntake) {
	rows, err := Config.DB.Model(&DB.WaterIntake{}).Where("user_id = ?", userId).Rows()
	if err != nil {
		return err, nil
	} else {
		var intakes []DB.WaterIntake
		for rows.Next() {
			var intake DB.WaterIntake
			err = Config.DB.ScanRows(rows, &intake)

			if err != nil {
				return err, nil
			}

			intakes = append(intakes, intake)
		}
		return nil, intakes
	}
}

func getTotalTodayWaterIntake(intakes []DB.WaterIntake) int {
	total := 0
	for _, intake := range intakes {
		total += intake.Amount
	}
	return total
}

func getSleepDuration(userId int) *Widgets.SleepSubWidget {
	var sleepTime DB.SleepTime
	err := Config.DB.Where("user_id = ?", userId).Last(&sleepTime).Error

	if err != nil {
		return nil
	} else {
		duration := sleepTime.To.Sub(sleepTime.From)
		return &Widgets.SleepSubWidget{
			Hours:   int(duration.Hours()),
			Minutes: int(duration.Minutes()) - int(duration.Hours())*60,
		}
	}
}

func mapDBIntakesToWidget(intakes []DB.WaterIntake) []Widgets.WaterIntake {
	var widgetIntakes []Widgets.WaterIntake
	for _, item := range intakes {
		widgetIntakes = append(widgetIntakes, Widgets.WaterIntake{
			TimeDiapason:   item.Time.String(),
			AmountInMillis: item.Amount,
		})
	}
	return widgetIntakes
}
