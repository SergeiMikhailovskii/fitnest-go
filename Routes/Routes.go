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
	}
	return r
}
