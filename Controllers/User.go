package Controllers

import (
	"TestProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginUser(c *gin.Context) {
	var userRequest Models.User
	var userResponse Models.User
	_ = c.BindJSON(&userRequest)
	err := Models.GetUserByID(&userResponse, userRequest.Login)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userResponse)
	}
}
