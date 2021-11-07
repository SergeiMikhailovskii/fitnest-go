package Controllers

import (
	"TestProject/Controllers/Handlers"
	"TestProject/Models/Base"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMainPage(c *gin.Context) {
	var responseStatusCode = -1
	var response Base.Response
	if Handlers.HasAuthUserCookie(c) {
		if Handlers.IsOnboardingFinished(c) {
			responseStatusCode = http.StatusOK
			response = Base.Response{
				Flow: "/onboarding",
			}
		} else {
			responseStatusCode = http.StatusOK
			response = Base.Response{
				Flow: "/onboarding",
			}
		}
	} else {
		responseStatusCode, response = Handlers.GenerateAuthUserToken(c)
	}
	c.JSON(responseStatusCode, response)
}
