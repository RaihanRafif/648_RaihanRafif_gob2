package router

import "github.com/gin-gonic/gin"

func StartApp() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/orders")
	{
		userRouter.POST("/")
	}

	return r
}
