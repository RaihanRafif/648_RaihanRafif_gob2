package routers

import (
	config "ass2/config"
	"ass2/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func StartServer() *gin.Engine {

	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	router := gin.Default()

	router.POST("/orders", inDB.CreateOrders)
	router.GET("/orders", inDB.GetOrders)
	router.PUT("/orders/:orderId", inDB.UpdateOrder)
	router.DELETE("/orders/:orderId", inDB.DeleteOrder)

	return router
}
