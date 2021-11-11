package Routes

import (
	"TestProject/Controllers/User"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("login", User.LoginUser)
		authGroup.POST("register", User.RegisterUser)
	}
}
