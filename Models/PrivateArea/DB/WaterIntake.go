package DB

import "time"

type WaterIntake struct {
	ID     int       `json:"id"`
	UserId int       `json:"user_id"`
	Time   time.Time `json:"time"`
	Amount int       `json:"amount"`
}

type ActivityTrackerSumQuery struct {
	TotalWaterIntake int
	TotalSteps       int
}

func (b *WaterIntake) TableName() string {
	return "water_intake"
}
