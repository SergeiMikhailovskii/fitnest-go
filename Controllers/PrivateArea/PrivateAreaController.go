package PrivateArea

import (
	"TestProject/Controllers/PrivateArea/ActivityTracker"
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

func GetActivityTrackerPage(c *gin.Context) {
	var response Base.Response
	data, err := ActivityTracker.GetActivityTrackerPage(c)

	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "private_area_activity_tracker",
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

func PinNotification(c *gin.Context) {
	var response Base.Response
	err := Notification.PinNotification(c)

	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "private_area_notifications_pin",
				Message: err.Error(),
			},
		}
	}
	response.Flow = "/private-area"

	c.JSON(http.StatusOK, response)
}

func DeleteNotification(c *gin.Context) {
	var response Base.Response
	err := Notification.DeleteNotification(c)

	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "private_area_notifications_delete",
				Message: err.Error(),
			},
		}
	}
	response.Flow = "/private-area"

	c.JSON(http.StatusOK, response)
}

func DeleteActivity(c *gin.Context) {
	var response Base.Response
	err := ActivityTracker.DeleteActivity(c)

	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "private_area_activity_delete",
				Message: err.Error(),
			},
		}
	}
	response.Flow = "/private-area"

	c.JSON(http.StatusOK, response)
}

func AddActivity(c *gin.Context) {
	var response Base.Response
	err := ActivityTracker.AddActivity(c)

	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "private_area_activity_add",
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
	Dashboard.GenerateCaloriesIntakeStub(userId)
	Dashboard.GenerateStepsStub(userId)
	Dashboard.GenerateWaterIntakeAimStub(userId)
	Dashboard.GenerateSleepTimeStub(userId)
	Dashboard.GenerateHeartRateStub(userId)
}
