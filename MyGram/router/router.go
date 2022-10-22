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
		photoRouter.PUT("/:photoId", middleware.UserAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.UserAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", middleware.UserAuthorization(), controllers.PostComment)
		commentRouter.GET("/", middleware.UserAuthorization(), controllers.GetComment)
		commentRouter.PUT("/:commentId", middleware.UserAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.UserAuthorization(), controllers.DeleteComment)
	}

	socialMediasRouter := r.Group("/socialmedias")
	{
		socialMediasRouter.Use(middleware.Authentication())
		socialMediasRouter.POST("/", middleware.UserAuthorization(), controllers.PostSocialMedia)
		socialMediasRouter.GET("/", middleware.UserAuthorization(), controllers.GetSocialMedia)
		socialMediasRouter.PUT("/:socialMediaId", middleware.UserAuthorization(), controllers.UpdateSocialMedia)
		socialMediasRouter.DELETE("/:socialMediaId", middleware.UserAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
