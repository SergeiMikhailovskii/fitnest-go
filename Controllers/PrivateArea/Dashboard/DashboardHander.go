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
	widgetsMap := make(map[string]interface{})

	widgetsMap["HEADER_WIDGET"] = getHeaderWidget(c)
	widgetsMap["BMI_WIDGET"] = getBMIWidget(c)
	widgetsMap["TODAY_TARGET_WIDGET"] = getTodayTargetWidget()
	widgetsMap["ACTIVITY_STATUS_WIDGET"] = getActivityStatusWidget(c)
	widgetsMap["LATEST_WORKOUT_WIDGET"] = getLatestWorkoutWidget()

	return &PrivateArea.Response{
		Widgets: widgetsMap,
	}, nil
}

func getHeaderWidget(c *gin.Context) Widgets.HeaderWidget {
	primaryRecord := Registration.GetPrimaryRegistrationRecord(c)
	userName := primaryRecord.FirstName + " " + primaryRecord.LastName

	userId, _ := Registration.GetUserId(c)

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

func getActivityStatusWidget(c *gin.Context) Widgets.ActivityStatusWidget {
	userId, _ := Registration.GetUserId(c)

	heartRate, date, err := getLastHeartRate(userId)

	var heartRateWidget *Widgets.HeartRateSubWidget

	if err == nil {
		heartRateWidget = &Widgets.HeartRateSubWidget{
			Rate: heartRate,
			Date: date,
		}
	} else {
		heartRateWidget = nil
	}

	return Widgets.ActivityStatusWidget{
		HeartRate: heartRateWidget,
		WaterIntake: Widgets.WaterIntakeSubWidget{
			Amount:   4,
			Progress: 0.5,
			Intakes: []Widgets.WaterIntake{
				{
					TimeDiapason:   "6am - 8am",
					AmountInMillis: 600,
				},
				{
					TimeDiapason:   "9am - 11am",
					AmountInMillis: 500,
				},
				{
					TimeDiapason:   "11am - 2pm",
					AmountInMillis: 1000,
				},
				{
					TimeDiapason:   "2pm - 4pm",
					AmountInMillis: 700,
				},
				{
					TimeDiapason:   "4pm - now",
					AmountInMillis: 900,
				},
			},
		},
		Sleep: Widgets.SleepSubWidget{
			Duration: Widgets.SleepDuration{
				Hours:   8,
				Minutes: 20,
			},
		},
		Calories: Widgets.CaloriesSubWidget{
			Consumed: 760,
			Left:     230,
		},
	}
}

func getLatestWorkoutWidget() Widgets.LatestWorkoutWidget {
	return Widgets.LatestWorkoutWidget{
		Workouts: []Widgets.Workout{
			{
				Name:     "Fullbody Workout",
				Calories: 180,
				Minutes:  20,
				Progress: 0.5,
			},
			{
				Name:     "Lowerbody Workout",
				Calories: 200,
				Minutes:  30,
				Progress: 0.4,
			},
			{
				Name:     "Ab Workout",
				Calories: 220,
				Minutes:  40,
				Progress: 0.3,
			},
		},
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
