package Models

import "github.com/golang-jwt/jwt"

type User struct {
	jwt.StandardClaims

	ID       int    `json:"id"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (b *User) TableName() string {
	return "user"
}
