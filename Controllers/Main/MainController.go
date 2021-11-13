package Main

import (
	"TestProject/Models/Base"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMainPage(c *gin.Context) {
	var responseStatusCode = -1
	var response Base.Response
	if HasAuthUserCookie(c) {
		if IsOnboardingFinished(c) {
			responseStatusCode = http.StatusOK
			response = Base.Response{
				Flow: Util.AfterOnboarding,
			}
		} else {
			responseStatusCode = http.StatusOK
			response = Base.Response{
				Flow: Util.Onboarding,
			}
		}
	} else {
		responseStatusCode, response = GenerateAuthUserToken(c)
	}
	c.JSON(responseStatusCode, response)
}
