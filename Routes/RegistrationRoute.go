package Routes

import (
	"TestProject/Controllers/Registration"
	"github.com/gin-gonic/gin"
)

func SetupRegistrationRoute(r *gin.Engine) {
	registrationGroup := r.Group("/registration")
	{
		registrationGroup.GET("", Registration.GetStep)
		registrationGroup.POST("", Registration.SubmitStep)
	}
}
