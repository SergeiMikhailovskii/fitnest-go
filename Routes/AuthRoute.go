package Routes

import (
	"TestProject/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("login", Controllers.LoginUser)
		authGroup.POST("register", Controllers.RegisterUser)
	}
}
