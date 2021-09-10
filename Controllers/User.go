package Controllers

import (
	"TestProject/Errors"
	"TestProject/Models"
	"TestProject/Models/Base"
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
		c.JSON(http.StatusOK, userResponse)
	}
}
