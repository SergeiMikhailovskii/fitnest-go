package Models

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (b *User) TableName() string {
	return "user"
}
