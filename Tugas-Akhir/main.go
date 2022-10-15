package main

import (
	"tugasakhir/database"
	"tugasakhir/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
