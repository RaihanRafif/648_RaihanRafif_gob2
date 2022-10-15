package router

import (
	"net/http"
	"tugasakhir/controllers"
	middleware "tugasakhir/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.Use(middleware.Authentication())
		userRouter.PUT("/", middleware.UserAuthorization(), controllers.UserUpdate)
		userRouter.DELETE("/", middleware.UserAuthorization(), controllers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", middleware.UserAuthorization(), controllers.PhotoUploader)
		r.StaticFS("/file", http.Dir("assets"))
		photoRouter.GET("/", middleware.UserAuthorization(), controllers.GetPhoto)
	}

	return r
}
