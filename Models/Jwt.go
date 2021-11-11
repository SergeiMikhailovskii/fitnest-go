package Models

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

func GenerateJwt(userId int) (string, error) {
	signingKey := []byte("8kkeN4jhL4F84qfw")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["userId"] = userId

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		_ = fmt.Errorf("something went wrong %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ParseJwt(token string) User {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("8kkeN4jhL4F84qfw"), nil
	})

	if err != nil {
		panic(err)
	}
	return User{
		ID: int(claims["userId"].(float64)),
	}
}
