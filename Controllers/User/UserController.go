package User

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GenerateToken(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, _ := strconv.Atoi(userIdStr)
	jwt, _ := Models.GenerateJwt(userId)
	cookie := Base.AuthUserCookie
	cookie.Value = jwt
	Util.SetDefaultCookie(c, cookie)
	c.JSON(http.StatusOK, jwt)
}
