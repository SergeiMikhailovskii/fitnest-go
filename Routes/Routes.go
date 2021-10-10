package Routes

import (
	"TestProject/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	authGroup := r.Group("/auth")
	{
		authGroup.POST("login", Controllers.LoginUser)
		authGroup.POST("register", Controllers.RegisterUser)
		authGroup.GET("generate-token", Controllers.GenerateToken)
	}
	return r
}
