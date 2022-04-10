package Main

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Models/Onboarding"
	"TestProject/Models/Registration"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HasAuthUserCookie(c *gin.Context) bool {
	_, err := c.Cookie(Base.AuthUserCookie.Name)
	return err == nil
}

func GenerateAuthUserToken(c *gin.Context) (int, Base.Response) {
	newUser := createNewUser()
	setAuthUserToken(newUser, c)
	return http.StatusOK, Base.Response{
		Flow: Util.Onboarding,
	}
}

func createNewUser() Models.User {
	newUser := Models.User{}
	_ = Models.CreateUser(&newUser)
	_ = Onboarding.CreateOnboardingDefaultRecord(newUser.ID)
	_ = Registration.CreatePrimaryRegistrationRecord(newUser.ID)
	return newUser
}

func setAuthUserToken(user Models.User, c *gin.Context) {
	token, _ := Models.GenerateJwt(user.ID)
	Base.AuthUserCookie.Value = token
	Util.SetDefaultCookie(c, Base.AuthUserCookie)
}
