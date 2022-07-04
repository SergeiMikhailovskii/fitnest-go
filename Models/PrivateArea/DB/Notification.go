package DB

import "time"

type Notification struct {
	ID       int       `json:"id"`
	UserId   int       `json:"user_id"`
	Text     string    `json:"text"`
	Date     time.Time `json:"date"`
	IsActive bool      `json:"is_active" gorm:"default:true"`
	IsPinned bool      `json:"is_pinned" gorm:"default:false"`
	Type     string    `json:"type" gorm:"default:DEFAULT"`
}

func (b *Notification) TableName() string {
	return "notification"
}
