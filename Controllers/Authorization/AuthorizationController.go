package Authorization

import (
	"TestProject/Models"
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
		Data: getLoginResponse,
		Flow: Util.Login,
	}

	c.JSON(http.StatusOK, response)
}

func LoginUser(c *gin.Context) {
	var request Authorization.GetLoginFields
	var response Base.Response
	_ = c.BindJSON(&request)
	err, userId := loginUser(request)

	if err != nil {
		response.Errors = []Base.Error{*err}
	} else {
		response.Flow = Util.Main
		jwt, _ := Models.GenerateJwt(*userId)
		cookie := Base.AuthUserCookie
		cookie.Value = jwt
		Util.SetDefaultCookie(c, cookie)
	}
	c.JSON(http.StatusOK, response)
}

func ForgetPassword(c *gin.Context) {
	var request Authorization.ForgetPasswordFields
	_ = c.BindJSON(&request)

	err, password := getPassword(*request.Login)

	if err == nil {
		sendForgetPasswordEmail(*request.Login, password)
	}
	c.JSON(http.StatusOK, Base.Response{})
}
