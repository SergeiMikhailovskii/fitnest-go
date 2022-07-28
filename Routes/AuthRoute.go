package Routes

import (
	"TestProject/Controllers/Authorization"
	"TestProject/Controllers/User"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.GET("token/:userId", User.GenerateToken)
		authGroup.GET("login", Authorization.GetLoginPage)
		authGroup.POST("login", Authorization.LoginUser)
		authGroup.POST("forget-password", Authorization.ForgetPassword)
	}
}
