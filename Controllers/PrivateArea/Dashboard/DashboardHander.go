package Dashboard

import (
	"TestProject/Controllers/Registration"
	"TestProject/Models/PrivateArea"
	"TestProject/Models/PrivateArea/Widgets"
	"github.com/gin-gonic/gin"
	"time"
)

func GetDashboardPage(c *gin.Context) (*PrivateArea.Response, error) {
	widgetsMap := make(map[string]interface{})

	widgetsMap["HEADER_WIDGET"] = getHeaderWidget(c)
	widgetsMap["BMI_WIDGET"] = getBMIWidget()
	widgetsMap["TODAY_TARGET_WIDGET"] = getTodayTargetWidget()
	widgetsMap["ACTIVITY_STATUS_WIDGET"] = getActivityStatusWidget()
	widgetsMap["LATEST_WORKOUT_WIDGET"] = getLatestWorkoutWidget()

	return &PrivateArea.Response{
		Widgets: widgetsMap,
	}, nil
}

func getHeaderWidget(c *gin.Context) Widgets.HeaderWidget {
	primaryRecord := Registration.GetPrimaryRegistrationRecord(c)
	userName := primaryRecord.FirstName + " " + primaryRecord.LastName

	return Widgets.HeaderWidget{
		Name:          userName,
		Notifications: 0,
	}
}

func getBMIWidget() Widgets.BMIWidget {
	return Widgets.BMIWidget{
		Index:  20.1,
		Result: "NORMAL_WEIGHT",
	}
}

func getTodayTargetWidget() Widgets.TodayTargetWidget {
	return Widgets.TodayTargetWidget{}
}

func getActivityStatusWidget() Widgets.ActivityStatusWidget {
	return Widgets.ActivityStatusWidget{
		HeartRate: Widgets.HeartRateSubWidget{
			Rate: 78,
			Date: time.Now(),
		},
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
