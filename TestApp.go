package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	addLoginRoutes(r.Group("/auth"))

	err := r.Run()
	if err != nil {
		return
	}
}
