package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/controllers"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/user")
	{
		userGroup.POST("/login", controllers.Login)
		userGroup.POST("/register", controllers.Register)
		userGroup.PUT("/update", middlewares.AuthMiddleware(), controllers.UpdateUser)
		userGroup.PUT("/password", middlewares.AuthMiddleware(), controllers.UpdatePassword)
		userGroup.GET("/", middlewares.AuthMiddleware(), controllers.GetUserInfo)
	}

	photoGroup := r.Group("/photo")
	{
		photoGroup.GET("/:fileName", controllers.ViewPhotoProfile)
		photoGroup.POST("/upload", middlewares.AuthMiddleware(), controllers.AddPhotoProfile)
		photoGroup.PUT("/update", middlewares.AuthMiddleware(), controllers.UpdatePhotoProfile)
		photoGroup.DELETE("/delete", middlewares.AuthMiddleware(), controllers.DeletePhotoProfile)
	}

	return r
}
