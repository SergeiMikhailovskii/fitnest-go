package Routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	SetupAuthRoute(r)
	SetupMainRoute(r)
	SetupOnboardingRoute(r)
	return r
}
