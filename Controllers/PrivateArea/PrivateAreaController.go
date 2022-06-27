package PrivateArea

import (
	"TestProject/Controllers/PrivateArea/Dashboard"
	"TestProject/Controllers/PrivateArea/Notification"
	"TestProject/Controllers/Registration"
	"TestProject/Models/Base"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDashboardPage(c *gin.Context) {
	var response Base.Response
	data, err := Dashboard.GetDashboardPage(c)

	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "private_area_dashboard",
				Message: err.Error(),
			},
		}
	}
	response.Data = &data
	response.Flow = "/private-area"

	c.JSON(http.StatusOK, response)
}

func GetNotificationsPage(c *gin.Context) {
	var response Base.Response
	data, err := Notification.GetNotificationsPage(c)

	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "private_area_notifications",
				Message: err.Error(),
			},
		}
	}
	response.Data = &data
	response.Flow = "/private-area"

	c.JSON(http.StatusOK, response)
}

func DeactivateNotifications(c *gin.Context) {
	var response Base.Response
	err := Notification.DeactivateNotifications(c)

	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "private_area_notifications_deactivate",
				Message: err.Error(),
			},
		}
	}
	response.Flow = "/private-area"

	c.JSON(http.StatusOK, response)
}

func GenerateDashboardStub(c *gin.Context) {
	userId, _ := Registration.GetUserId(c)

	Dashboard.GenerateNotificationsStub(userId)
	Dashboard.GenerateWorkoutsStub()
	Dashboard.GenerateUserWorkoutsStub(userId)
	Dashboard.GenerateWaterIntakeStub(userId)
	Dashboard.GenerateWaterIntakeStub(userId)
	Dashboard.GenerateCaloriesIntakeStub(userId)
	Dashboard.GenerateWaterIntakeAimStub(userId)
	Dashboard.GenerateSleepTimeStub(userId)
	Dashboard.GenerateHeartRateStub(userId)
}
