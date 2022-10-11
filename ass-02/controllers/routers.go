package controllers

import (
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", createOrder("raihan"))

	return router
}
