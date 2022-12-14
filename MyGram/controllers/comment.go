package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"tugasakhir/database"
	"tugasakhir/helpers"
	"tugasakhir/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type GetCommentCustomPhoto struct {
	id       int
	Title    string
	Caption  string
	PhotoUrl string
	UserId   int
}

type GetCommentCustomUser struct {
	Id       int
	Email    string
	Username string
}

type GetCommentCustom struct {
	Id        int
	Message   string
	PhotoId   int
	UserId    int
	CreatedAt string
	UpdatedAt string
	User      GetCommentCustomUser
	Photo     GetCommentCustomPhoto
}

func PostComment(c *gin.Context) {
	db := database.GetDB()
	dt := time.Now()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.CreatedAt = dt.Format("01-02-2006 15:04:05")
	Comment.UpdatedAt = dt.Format("01-02-2006 15:04:05")
	Comment.UserID = userID

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

	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	db.Find(&Comment, "user_id = ?", userID)

	var results []GetCommentCustom

	for i := 0; i < len(Comment); i++ {
		Photo := models.Photo{}

		db.Find(&User, "id = ?", Comment[i].PhotoID)
		customUser := GetCommentCustomUser{
			Id:       int(User.ID),
			Email:    User.Email,
			Username: User.UserName,
		}

		db.Find(&Photo, "id = ?", Comment[i].PhotoID)
		customPhoto := GetCommentCustomPhoto{
			int(Photo.ID),
			Photo.Title,
			Photo.Caption,
			Photo.PhotoUrl,
			int(Photo.UserID),
		}
		tes := GetCommentCustom{
			int(Comment[i].ID),
			Comment[i].Message,
			int(Comment[i].PhotoID),
			int(Comment[i].UserID),
			Comment[i].UpdatedAt,
			Comment[i].CreatedAt,
			customUser,
			customPhoto,
		}

		results = append(results, tes)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": results,
	})
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	Comment := models.Comment{}
	Photo := models.Photo{}
	commentId, _ := strconv.Atoi(c.Param("commentId"))
	dt := time.Now()

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	// err := db.Find(&Comment).Where("id = ?", commentId).Error
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"err":     "Bad Request",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// err = db.Find(&Photo, "id = ?", Comment.PhotoID).Error
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"err":     "Bad Request",
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	Comment.UpdatedAt = dt.Format("01-02-2006 15:04:05")
	err := db.Model(&Comment).Where("id = ?", commentId).Where("user_id = ?", userID).Updates(models.Comment{Message: Comment.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	db.Where("id = ?", commentId).Find(&Comment)
	db.Where("id = ?", Comment.PhotoID).Find(&Photo)
	fmt.Println("Comment", Comment)
	c.JSON(http.StatusOK, gin.H{
		"id":         commentId,
		"message":    Comment.Message,
		"caption":    Photo.Caption,
		"photo_url":  Photo.PhotoUrl,
		"user_id":    userID,
		"updated_at": dt.Format("01-02-2006 15:04:05"),
	})
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
