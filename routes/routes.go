package routes

import (
	"medigo-be/controllers"
	"medigo-be/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/me", middleware.AuthMiddleware(), controllers.Me)
	}

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		obat := protected.Group("/api/obat")
		{
			obat.GET("", controllers.GetAllObat)
			obat.GET("/:id", controllers.GetObatByID)
			obat.POST("", controllers.CreateObat)
			obat.PUT("/:id", controllers.UpdateObat)
			obat.DELETE("/:id", controllers.DeleteObat)
		}
	}

}
