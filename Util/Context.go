package Util

import (
	"TestProject/Models/Base"
	"github.com/gin-gonic/gin"
)

func SetDefaultCookie(c *gin.Context, cookie Base.Cookie) {
	c.SetCookie(cookie.Name, cookie.Value, 0, "/", "", false, true)
}
