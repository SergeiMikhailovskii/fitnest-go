package DB

import "time"

type SleepTime struct {
	ID     int       `json:"id"`
	UserId int       `json:"user_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}

func (b *HeartRate) SleepTime() string {
	return "sleep_time"
}
