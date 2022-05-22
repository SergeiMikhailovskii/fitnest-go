package PrivateArea

import (
	"TestProject/Controllers/PrivateArea/Dashboard"
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
}
