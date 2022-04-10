package Main

import (
	"TestProject/Controllers/Onboarding"
	"TestProject/Controllers/Registration"
	"TestProject/Models/Base"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMainPage(c *gin.Context) {
	var responseStatusCode = -1
	var response Base.Response
	if HasAuthUserCookie(c) {
		if !Onboarding.IsOnboardingFinished(c) {
			responseStatusCode = http.StatusOK
			response = Base.Response{
				Flow: Util.Onboarding,
			}
		} else if !Registration.IsRegistrationFinished(c) {
			responseStatusCode = http.StatusOK
			response = Base.Response{
				Flow: Util.Registration,
			}
		} else {
			responseStatusCode = http.StatusOK
			response = Base.Response{
				Flow: Util.Main,
			}
		}
	} else {
		responseStatusCode, response = GenerateAuthUserToken(c)
	}
	c.JSON(responseStatusCode, response)
}
