package Routes

import (
	"TestProject/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupMainRoute(r *gin.Engine) {
	authGroup := r.Group("/main")
	{
		authGroup.GET("", Controllers.GenerateToken)
	}
}
