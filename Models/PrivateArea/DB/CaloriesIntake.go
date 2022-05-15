package DB

import "time"

type CaloriesIntake struct {
	ID     int       `json:"id"`
	UserId int       `json:"user_id"`
	Time   time.Time `json:"time"`
	Amount int       `json:"amount"`
}

func (b *CaloriesIntake) TableName() string {
	return "calories_intake"
}
