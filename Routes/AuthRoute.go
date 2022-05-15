package Routes

import (
	"TestProject/Controllers/User"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.GET("token/:userId", User.GenerateToken)
	}
}
