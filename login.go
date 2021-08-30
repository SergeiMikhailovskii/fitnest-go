package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addLoginRoutes(rg *gin.RouterGroup) {
	login := rg.Group("/login")
	login.POST("/", func(context *gin.Context) {
		handleLoginRequest(context)
	})
}

func handleLoginRequest(context *gin.Context) {
	var request LoginData
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if request.Login == "Sergei" && request.Password == "12345" {
		context.JSON(http.StatusOK, gin.H{})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"error": errors.New("user not found").Error(),
		})
	}

}

type LoginData struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
