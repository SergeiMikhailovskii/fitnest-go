package User

import (
	"TestProject/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GenerateToken(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)
	jwt, _ := Models.GenerateJwt(userId)
	c.JSON(http.StatusOK, jwt)
}
