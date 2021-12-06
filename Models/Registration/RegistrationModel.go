package Registration

import "time"

type PrimaryInfo struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Sex       string    `json:"sex"`
	BirthDate time.Time `json:"birth_date"`
}

func (b *PrimaryInfo) TableName() string {
	return "registration_primary_info"
}
