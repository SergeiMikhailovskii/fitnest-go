package DB

import "time"

type Notification struct {
	ID       int       `json:"id"`
	UserId   int       `json:"user_id"`
	ImageURL string    `json:"image_url"`
	Text     string    `json:"text"`
	Date     time.Time `json:"date"`
	IsActive bool      `json:"is_active" gorm:"default:true"`
}

func (b *Notification) TableName() string {
	return "notification"
}
