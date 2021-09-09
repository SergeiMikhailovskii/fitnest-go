package Controllers

import (
	"TestProject/Models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

func LoginUser(c *gin.Context) {
	var userRequest Models.User
	var userResponse Models.User
	_ = c.BindJSON(&userRequest)
	err := Models.GetUserByID(&userResponse, userRequest)

	if err.Error() == "user not found" {
		c.JSON(http.StatusOK, render.JSON{
			Data: err.Error(),
		})
	} else if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userResponse)
	}
}
