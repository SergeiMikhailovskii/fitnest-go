package Handlers

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckAuthUserCookie checks if user has AuthUser cookie and if it exists returns true else setting it,
// sending response with status unauthorized and returns false
func CheckAuthUserCookie(c *gin.Context) bool {
	_, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		newUser := createNewUser()
		setAuthUserToken(newUser, c)
		c.JSON(http.StatusUnauthorized, Base.Response{})
	}
	return err == nil
}

func createNewUser() Models.User {
	newUser := Models.User{}
	_ = Models.CreateUser(&newUser)
	return newUser
}

func setAuthUserToken(user Models.User, c *gin.Context) {
	token, _ := Util.GenerateJwt(user.ID)
	Base.AuthUserCookie.Value = token
	Util.SetDefaultCookie(c, Base.AuthUserCookie)
}
