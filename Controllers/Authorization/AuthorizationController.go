package Authorization

import (
	"TestProject/Models/Authorization"
	"TestProject/Models/Base"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLoginPage(c *gin.Context) {
	fields := getLoginFields()
	validationSchema := getLoginValidationSchema()

	getLoginResponse := Authorization.GetLoginResponse{
		Fields:           fields,
		ValidationSchema: validationSchema,
	}

	response := Base.Response{
		Data:   getLoginResponse,
		Errors: nil,
		Flow:   Util.Login,
	}

	c.JSON(http.StatusOK, response)
}
