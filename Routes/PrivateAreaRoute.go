package Routes

import (
	"TestProject/Controllers/PrivateArea"
	"github.com/gin-gonic/gin"
)

func SetupPrivateAreaRoute(r *gin.Engine) {
	privateAreaGroup := r.Group("/private-area")
	{
		privateAreaGroup.GET("/dashboard", PrivateArea.GetDashboardPage)
		privateAreaGroup.GET("/dashboard/generate-stubs", PrivateArea.GenerateDashboardStub)
	}
}
