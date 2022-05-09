package PrivateArea

import (
	"TestProject/Controllers/PrivateArea/Dashboard"
	"TestProject/Models/Base"
	"TestProject/Models/PrivateArea"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPage(c *gin.Context) {
	var response Base.Response
	var data *PrivateArea.Response
	var err error

	page := c.Param("page")

	switch page {
	case "DASHBOARD":
		data, err = Dashboard.GetDashboardPage(c)
	}
	if err != nil {
		response.Errors = []Base.Error{
			{
				Field:   "registration",
				Message: err.Error(),
			},
		}
	}
	response.Data = &data
	c.JSON(http.StatusOK, response)
}
