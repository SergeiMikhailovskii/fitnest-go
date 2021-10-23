package Controllers

import (
	"TestProject/Controllers/Handlers"
	"TestProject/Models/Base"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMainPage(c *gin.Context) {
	if Handlers.CheckAuthUserCookie(c) {
		c.JSON(http.StatusOK, Base.Response{})
	}
}
