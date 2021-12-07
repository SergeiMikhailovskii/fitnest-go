package Registration

import (
	"TestProject/Models/Base"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
)

func GetStep(c *gin.Context) {
	var responseStatusCode = -1
	var response Base.Response
	response.Flow = Util.Registration
	data, err := getRegistrationStep(c)
	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "registration",
				Message: err.Error(),
			},
		}
	}
	response.Data = &data
	c.JSON(responseStatusCode, response)
}

func SubmitStep(c *gin.Context) {

}
