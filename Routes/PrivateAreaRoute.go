package Routes

import (
	"TestProject/Controllers/PrivateArea"
	"github.com/gin-gonic/gin"
)

func SetupPrivateAreaRoute(r *gin.Engine) {
	privateAreaGroup := r.Group("/private-area")
	{
		privateAreaGroup.GET("/:page", PrivateArea.GetPage)
	}
}
