package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yardan-org/final-task-pbi-rakamin-fullstack-yardan/app/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.Register)
	}

	return r
}
