package controllers

import (
	"encoding/json"
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
	Photo.UserID = userID

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

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	result := db.Find(&Photo, "user_id = ?", userID)

	p, _ := json.Marshal(result)
	err := json.Unmarshal(p, &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("PPPPP", result)

	c.JSON(http.StatusOK, gin.H{
		"token": err,
	})

}
