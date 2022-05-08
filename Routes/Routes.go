package Routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	SetupAuthRoute(r)
	SetupRegistrationRoute(r)
	SetupFlowRoute(r)
	SetupOnboardingRoute(r)
	SetupPrivateAreaRoute(r)
	return r
}
