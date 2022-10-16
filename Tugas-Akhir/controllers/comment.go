package controllers

import (
	"net/http"
	"strconv"
	"tugasakhir/database"
	"tugasakhir/helpers"
	"tugasakhir/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	db := database.GetDB()
	// userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	// userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := []models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	db.Find(&Comment, "user_id = ?", userID).Joins("INNER JOIN users ON comments.user_id=users.id")
	// db.Model(&Comment).Select("comments.user_id, comments.photo_id").Joins("left join photos on photos.user_id = %s", userID)

	c.JSON(http.StatusOK, gin.H{
		"result": Comment,
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	Comment := models.Comment{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Model(&Comment).Where("id = ?", commentId).Where("user_id = ?", userID).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	Comment := models.Comment{}

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	err := db.Where("id = ?", commentId).Where("user_id =?", userID).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Comment has been sucessfully deleted",
	})
}
