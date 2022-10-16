package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"tugasakhir/database"
	"tugasakhir/helpers"
	"tugasakhir/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoUploader(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	dt := time.Now()
	currentTime := dt.Format("01-02-2006")
	// timeInSecond := dt.Format("15:04:05")

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	// GET Format File
	sourceFile := header.Filename
	splitedFile := strings.Split(sourceFile, ".")
	formatFile := splitedFile[1]

	filename := strconv.FormatUint(uint64(userID), 10) + "_" + currentTime + "_" + "." + formatFile
	out, err := os.Create("assets/" + filename)
	if err != nil {
		log.Fatal(err)
	}

	Photo.PhotoUrl = filename
	// Photo.UserID = userID

	err = db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, Photo)
}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := []models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	db.Find(&Photo, "user_id = ?", userID)

	c.JSON(http.StatusOK, gin.H{
		"result": Photo,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	Photo := models.Photo{}
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	err := db.Model(&Photo).Where("id = ?", photoId).Where("user_id = ?", userID).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	userID := uint(userData["id"].(float64))
	photoId, _ := strconv.Atoi(c.Param("photoId"))
	Photo := models.Photo{}

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	db.First(&Photo, photoId)
	// fmt.Println(Photo.PhotoUrl)
	e := os.Remove("assets/" + Photo.PhotoUrl)
	if e != nil {
		log.Fatal(e)
	}

	err := db.Where("id = ?", photoId).Where("user_id =?", userID).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your Photo has been sucessfully deleted",
	})
}
