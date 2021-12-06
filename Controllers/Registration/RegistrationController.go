package Registration

import (
	"TestProject/Models/Base"
	"TestProject/Models/Onboarding"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
)

func GetStep(c *gin.Context) {
	var responseStatusCode = -1
	var response Base.Response
	response.Flow = Util.Registration
	step, err := getRegistrationStep(c)
	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "registration",
				Message: err.Error(),
			},
		}
	}
	response.Data = Onboarding.Response{Step: step}
	c.JSON(responseStatusCode, response)
}

func SubmitStep(c *gin.Context) {

}
