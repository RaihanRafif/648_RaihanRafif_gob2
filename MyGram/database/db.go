package database

import (
	"fmt"
	"log"

	"tugasakhir/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "captainrr12"
	dbPort   = "5432"
	dbName   = "MyGram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	fmt.Println("sukses koneksi ke database")
	db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.SocialMedia{}, models.Photo{})
}

func GetDB() *gorm.DB {
	return db
}
