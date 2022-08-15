package DB

import "time"

type Steps struct {
	ID     int       `json:"id"`
	UserId int       `json:"user_id"`
	Time   time.Time `json:"time"`
	Amount int       `json:"amount"`
}

func (b *Steps) TableName() string {
	return "steps"
}
