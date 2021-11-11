package User

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginUser(c *gin.Context) {
	var userRequest Models.User
	var userResponse Models.User
	_ = c.BindJSON(&userRequest)
	err := Models.GetUserByCreds(&userResponse, userRequest)

	if errors.Is(err, Util.UserNotFound) {
		c.JSON(http.StatusOK, Base.Response{
			Errors: []Base.Error{{
				Field:   "login",
				Message: err.Error()},
			},
			Data: nil,
		})
	} else if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		authUserToken, _ := Models.GenerateJwt(userResponse.ID)
		http.SetCookie(c.Writer, &http.Cookie{
			Name:  "AuthUser",
			Value: authUserToken,
		})

		c.JSON(http.StatusOK, Base.Response{
			Errors: nil,
			Data:   struct{}{},
		})
	}
}

func RegisterUser(c *gin.Context) {
	var userRequest Models.User
	_ = c.BindJSON(&userRequest)

	err := Models.CheckUserExistsByLogin(userRequest)

	if errors.Is(err, Util.UserExists) {
		c.JSON(http.StatusOK, Base.Response{
			Errors: []Base.Error{{
				Field:   "login",
				Message: err.Error()},
			},
			Data: nil,
		})
		return
	} else if errors.Is(err, Util.UserNotFound) {
		Models.CreateUser(&userRequest)
		c.JSON(http.StatusOK, Base.Response{
			Errors: nil,
			Data:   nil,
		})
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, Base.Response{
			Errors: []Base.Error{{
				Field:   "login",
				Message: err.Error()},
			},
			Data: nil,
		})
	}
}