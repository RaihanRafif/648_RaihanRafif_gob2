package controllers

import (
	"live/database"
	"live/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRehan(c *gin.Context) {
	db := database.GetDB()

	Punyaku := models.Punyaku{}

	db.Find(&Punyaku)
	// db.Model(&Punyaku).Select("comments.user_id, comments.photo_id").Joins("left join photos on photos.user_id = %s", userID)

	c.JSON(http.StatusOK, gin.H{
		"result": Punyaku,
	})
}

func GetRehanConditional(c *gin.Context) {
	db := database.GetDB()
	paramku := c.Param("nameId")

	Punyaku := models.Punyaku{}

	// if contentType == appJson {
	// 	c.ShouldBindJSON(&CustomRespons)
	// } else {
	// 	c.ShouldBind(&CustomRespons)
	// }
	db.Find(&Punyaku)

	if paramku == "rehan" {
		tes := models.CustomRespons{
			201,
			true,
			Punyaku,
		}
		c.JSON(http.StatusOK, gin.H{
			"result": tes,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message code": 404,
			"status":       false,
		})
	}

	// // db.Model(&CustomRespons).Select("comments.user_id, comments.photo_id").Joins("left join photos on photos.user_id = %s", userID)

	// c.JSON(http.StatusOK, gin.H{
	// 	"result": CustomRespons,
	// })
}
