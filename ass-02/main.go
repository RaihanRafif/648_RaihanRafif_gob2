package main

import (
	"ass-2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()

	router.GET("/person/:id", controllers.CreateItem())
	router.Run(":3000")
}
