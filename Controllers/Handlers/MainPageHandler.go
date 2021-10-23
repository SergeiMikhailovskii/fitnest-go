package Handlers

import (
	"TestProject/Models/Base"
	"TestProject/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAuthUserCookie(c *gin.Context) bool {
	authUserCookie, err := c.Cookie(Base.AuthUserCookie.Name)
	Base.AuthUserCookie.Value = authUserCookie
	fmt.Println(authUserCookie)
	if err != nil {
		fmt.Println(err)
		Util.SetDefaultCookie(c, Base.AuthUserCookie)
		c.JSON(http.StatusUnauthorized, Base.Response{})
	}
	return Util.IsNotEmpty(authUserCookie)
}
