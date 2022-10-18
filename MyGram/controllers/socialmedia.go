package controllers

import (
	"net/http"
	"strconv"
	"time"
	"tugasakhir/database"
	"tugasakhir/helpers"
	"tugasakhir/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type GetSosmedCustomUser struct {
	Id              int
	UserName        string
	ProfileImageUrl string
}

type GetSosmedCustom struct {
	Id             int
	Name           string
	SocialMediaUrl string
	UserId         int
	CreatedAt      string
	UpdatedAt      string
	User           GetSosmedCustomUser
}

func PostSocialMedia(c *gin.Context) {
	db := database.GetDB()
	dt := time.Now()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.CreatedAt = dt.Format("01-02-2006 15:04:05")
	SocialMedia.UpdatedAt = dt.Format("01-02-2006 15:04:05")

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SocialMedia)
}

func GetSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := []models.SocialMedia{}
	Photo := models.Photo{}
	User := models.User{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	var results []GetSosmedCustom

	db.Find(&SocialMedia)

	db.Find(&Photo, "user_id = ?", userID)
	db.Find(&User, "id = ?", userID)
	customUser := GetSosmedCustomUser{
		Id:              int(User.ID),
		UserName:        User.UserName,
		ProfileImageUrl: Photo.PhotoUrl,
	}

	for i := 0; i < len(SocialMedia); i++ {
		tes := GetSosmedCustom{
			int(SocialMedia[i].ID),
			SocialMedia[i].Name,
			SocialMedia[i].SocialMediaUrl,
			int(SocialMedia[i].UserID),
			SocialMedia[i].CreatedAt,
			SocialMedia[i].UpdatedAt,
			customUser,
		}

		results = append(results, tes)
	}

	c.JSON(http.StatusOK, gin.H{
		"social_medias": results,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	SocialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Where("user_id = ?", userID).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	socialMediaId, _ := strconv.Atoi(c.Param("socialMediaId"))
	SocialMedia := models.SocialMedia{}

	if contentType == appJson {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	err := db.Where("id = ?", socialMediaId).Where("user_id =?", userID).Delete(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your SocialMedia has been sucessfully deleted",
	})
}
