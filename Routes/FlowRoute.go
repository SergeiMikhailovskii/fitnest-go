package Routes

import (
	"TestProject/Controllers/Registration/Flow"
	"github.com/gin-gonic/gin"
)

func SetupFlowRoute(r *gin.Engine) {
	authGroup := r.Group("/flow")
	{
		authGroup.GET("", Flow.GetFlow)
	}
}
