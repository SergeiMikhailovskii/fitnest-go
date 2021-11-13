package Main

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Models/Onboarding"
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
	return http.StatusUnauthorized, Base.Response{}
}

func IsOnboardingFinished(c *gin.Context) bool {
	cookie, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		panic(err)
	}

	Models.ParseJwt(cookie)

	return false
}

func GetOnboardingStep(c *gin.Context) string {
	cookie, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		panic(err)
	}

	user := Models.ParseJwt(cookie)

	onboardingRecord := Onboarding.Onboarding{}
	_ = Onboarding.GetRecordByUserId(user.ID, &onboardingRecord)

	if !onboardingRecord.FirstStepComplete {
		return "step1"
	} else if onboardingRecord.FirstStepComplete {
		return "step2"
	} else if onboardingRecord.SecondStepComplete {
		return "step3"
	} else if onboardingRecord.ThirdStepComplete {
		return "after"
	} else {
		return "undefined"
	}
}

func createNewUser() Models.User {
	newUser := Models.User{}
	_ = Models.CreateUser(&newUser)
	_ = Onboarding.CreateDefaultRecord(newUser.ID)
	return newUser
}

func setAuthUserToken(user Models.User, c *gin.Context) {
	token, _ := Models.GenerateJwt(user.ID)
	Base.AuthUserCookie.Value = token
	Util.SetDefaultCookie(c, Base.AuthUserCookie)
}
