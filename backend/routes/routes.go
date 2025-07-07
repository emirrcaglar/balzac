package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)

		// Protected routes (require auth)
		auth := api.Use(middleware.JWTAuth())
		{
			auth.GET("/profile", controllers.GetProfile)
		}
	}
}
