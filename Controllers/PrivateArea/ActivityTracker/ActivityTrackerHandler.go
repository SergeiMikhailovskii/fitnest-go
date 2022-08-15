package ActivityTracker

import (
	"TestProject/Config"
	"TestProject/Controllers/Registration"
	"TestProject/Models/PrivateArea"
	"TestProject/Models/PrivateArea/DB"
	"TestProject/Models/PrivateArea/Widgets"
	"github.com/gin-gonic/gin"
	"time"
)

func GetActivityTrackerPage(c *gin.Context) (*PrivateArea.Response, error) {
	userId, _ := Registration.GetUserId(c)

	widgetsMap := make(map[string]interface{})

	widgetsMap[PrivateArea.TodayTargetWidget] = getTodayTargetWidget(userId)
	widgetsMap[PrivateArea.ActivityProgressWidget] = getActivityProgressWidget(userId)
	widgetsMap[PrivateArea.LatestActivityWidget] = getLatestActivityWidget(userId)

	return &PrivateArea.Response{
		Widgets: widgetsMap,
	}, nil
}

func getActivityProgressWidget(userId int) *Widgets.ActivityProgressWidget {

	return &Widgets.ActivityProgressWidget{}
}

func getTodayTargetWidget(userId int) *Widgets.TodayTargetWidget {
	var result DB.ActivityTrackerSumQuery

	now := time.Now()
	dayStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	dayEnd := dayStart.Add(time.Hour * 24)

	_ = Config.DB.Model(&DB.WaterIntake{}).
		Select("sum(amount) as total_water_intake").
		Where("time BETWEEN ? AND ? AND user_id = ?", dayStart, dayEnd, userId).
		Find(&result)

	waterIntake := result.TotalWaterIntake

	_ = Config.DB.Model(&DB.Steps{}).
		Select("sum(amount) as total_steps").
		Where("time BETWEEN ? AND ? AND user_id = ?", dayStart, dayEnd, userId).
		Find(&result)

	steps := result.TotalSteps

	return &Widgets.TodayTargetWidget{
		WaterIntake: waterIntake,
		Steps:       steps,
	}
}

func getLatestActivityWidget(userId int) *Widgets.LatestActivityWidget {
	return &Widgets.LatestActivityWidget{}
}
