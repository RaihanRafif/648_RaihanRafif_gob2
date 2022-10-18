package main

import (
	"live/database"
	"live/routers"
)

func main() {
	database.StartDB()
	r := routers.StartApp()
	r.Run(":8080")
}
