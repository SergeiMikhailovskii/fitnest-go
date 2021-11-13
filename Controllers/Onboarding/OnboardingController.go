package Onboarding

import (
	"TestProject/Models/Base"
	"TestProject/Models/Onboarding"
	"github.com/gin-gonic/gin"
)

func GetStep(c *gin.Context) {
	var responseStatusCode = -1
	var response Base.Response
	response.Flow = "/onboarding"
	step := getOnboardingStep(c)
	response.Data = Onboarding.Response{Step: step}
	c.JSON(responseStatusCode, response)
}

func SubmitStep(c *gin.Context) {
	var responseStatusCode = -1
	var response Base.Response
	response.Flow = "/onboarding"
	err := submitOnboardingStep(c)
	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "onboarding",
				Message: err.Error(),
			},
		}
	}
	c.JSON(responseStatusCode, response)
}
