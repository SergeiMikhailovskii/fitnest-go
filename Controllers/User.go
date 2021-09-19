package Controllers

import (
	"TestProject/Errors"
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
	err := Models.GetUserByID(&userResponse, userRequest)

	if errors.Is(err, Errors.UserNotFound) {
		c.JSON(http.StatusOK, Base.Response{
			Errors: []string{err.Error()},
			Data:   nil,
		})
	} else if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		authUserToken, _ := Util.GenerateJwt(userResponse.ID)
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
	err := Models.CreateUser(&userRequest)

	if errors.Is(err, Errors.UserNotFound) {
		c.JSON(http.StatusOK, Base.Response{
			Errors: []string{err.Error()},
			Data:   nil,
		})
	} else if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Base.Response{
			Errors: nil,
			Data:   nil,
		})
	}
}
