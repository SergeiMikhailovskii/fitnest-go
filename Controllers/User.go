package Controllers

import (
	"TestProject/Errors"
	"TestProject/Models"
	"TestProject/Models/Base"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		http.SetCookie(c.Writer, &http.Cookie{
			Name:   "userId",
			Value:  strconv.Itoa(userResponse.ID),
			Domain: c.Request.RequestURI,
			Path:   "/",
		})
		c.JSON(http.StatusOK, userResponse)
	}
}
