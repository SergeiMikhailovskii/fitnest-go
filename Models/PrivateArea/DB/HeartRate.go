package DB

import "time"

type HeartRate struct {
	ID     int       `json:"id"`
	UserId int       `json:"user_id"`
	Rate   int       `json:"rate"`
	Date   time.Time `json:"date"`
}

func (b *HeartRate) TableName() string {
	return "heart_rate"
}
