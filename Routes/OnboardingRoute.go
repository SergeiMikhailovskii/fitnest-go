package Routes

import (
	"TestProject/Controllers/Onboarding"
	"github.com/gin-gonic/gin"
)

func SetupOnboardingRoute(r *gin.Engine) {
	authGroup := r.Group("/onboarding")
	{
		authGroup.GET("", Onboarding.GetStep)
		authGroup.POST("", Onboarding.SubmitStep)
	}
}
