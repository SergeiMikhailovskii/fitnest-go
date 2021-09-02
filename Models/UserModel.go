package Models

type User struct {
	ID       int
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (b *User) TableName() string {
	return "user"
}
