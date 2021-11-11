package Routes

import (
	"TestProject/Controllers/Main"
	"github.com/gin-gonic/gin"
)

func SetupMainRoute(r *gin.Engine) {
	authGroup := r.Group("/main")
	{
		authGroup.GET("", Main.GetMainPage)
	}
}
