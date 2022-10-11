package main

import (
	"secureApp/database"
	"secureApp/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
