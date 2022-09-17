package Routes

import (
	"TestProject/Controllers/PrivateArea"
	"github.com/gin-gonic/gin"
)

func SetupPrivateAreaRoute(r *gin.Engine) {
	privateAreaGroup := r.Group("/private-area")
	{
		privateAreaGroup.GET("/dashboard", PrivateArea.GetDashboardPage)
		privateAreaGroup.GET("/notifications", PrivateArea.GetNotificationsPage)
		privateAreaGroup.GET("/activity-tracker", PrivateArea.GetActivityTrackerPage)
		privateAreaGroup.GET("/dashboard/generate-stubs", PrivateArea.GenerateDashboardStub)

		privateAreaGroup.POST("/notifications/deactivate", PrivateArea.DeactivateNotifications)
		privateAreaGroup.POST("/notifications/pin", PrivateArea.PinNotification)
		privateAreaGroup.POST("/notifications/delete", PrivateArea.DeleteNotification)

		privateAreaGroup.POST("/activity-tracker/delete-activity", PrivateArea.DeleteActivity)
		privateAreaGroup.POST("/activity-tracker/add-activity", PrivateArea.AddActivity)
	}
}
