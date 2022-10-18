package routers

import (
	"live/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	r.GET("/rehan", controllers.GetRehan)

	r.GET("/data/:nameId", controllers.GetRehanConditional)

	return r
}
