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
	now := time.Now()
	periodEnd := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	periodStart := periodEnd.Add(-7 * time.Hour * 24)

	rows, err := Config.DB.Model(&DB.Steps{}).
		Select("date_trunc('day', time) AS date, SUM(amount) AS total").
		Where("time BETWEEN ? AND ? AND user_id = ?", periodStart, periodEnd, userId).
		Group("date").
		Rows()

	if err != nil {
		panic(err)
		return nil
	}

	var dbProgresses []Widgets.ActivityProgressItem
	var activityProgresses []Widgets.ActivityProgressItem

	for rows.Next() {
		var activityProgress DB.ActivityProgressQuery
		err = Config.DB.ScanRows(rows, &activityProgress)

		activityProgressWidget := Widgets.ActivityProgressItem{
			Date:  activityProgress.Date.Format("2006-01-02"),
			Total: activityProgress.Total,
		}

		if err != nil {
			return nil
		}

		dbProgresses = append(dbProgresses, activityProgressWidget)
	}

	for i := 0; i < 7; i++ {
		isDateFromDB := false
		currentDate := time.Date(periodStart.Year(), periodStart.Month(), periodStart.Day()+i, 0, 0, 0, 0, now.Location()).
			Format("2006-01-02")

		for _, item := range dbProgresses {
			if item.Date == currentDate {
				activityProgresses = append(activityProgresses, Widgets.ActivityProgressItem{Date: item.Date, Total: item.Total})
				isDateFromDB = true
			}
		}

		if !isDateFromDB {
			activityProgresses = append(activityProgresses, Widgets.ActivityProgressItem{Date: currentDate, Total: 0})
		}
	}

	return &Widgets.ActivityProgressWidget{
		Progresses: activityProgresses,
	}
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
