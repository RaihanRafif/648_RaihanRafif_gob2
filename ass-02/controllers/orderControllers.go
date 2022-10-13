package controllers

import (
	"ass-2/database"
	"ass-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDB()
	Order := models.Order{}

	err := db.Debug().Create(&Order).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"orderedAt":    User.ID,
		"customerName": User.Email,
		"items":        User.FullName,
	})
}
