package Handlers

import (
	"TestProject/Models/Base"
	"TestProject/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckAuthUserCookie checks if user has AuthUser cookie and if it exists returns true else setting it,
// sending response with status unauthorized and returns false
func CheckAuthUserCookie(c *gin.Context) bool {
	_, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		fmt.Println(err)
		Base.AuthUserCookie.Value = "test"
		Util.SetDefaultCookie(c, Base.AuthUserCookie)
		c.JSON(http.StatusUnauthorized, Base.Response{})
	}
	return err != nil
}
