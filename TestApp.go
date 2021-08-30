package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/login", func(context *gin.Context) {
		var request LoginData
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(request.Login)
		fmt.Println(request.Password)
		context.JSON(200, gin.H{
			"test": "it",
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}

type LoginData struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
