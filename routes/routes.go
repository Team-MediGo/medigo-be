package routes

import (
	"medigo-be/controllers"
	"medigo-be/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	// AUTH
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
		auth.GET("/me", middleware.AuthMiddleware(), controllers.Me)
	}

	// PROTECTED ROUTES
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// OBAT
		obat := protected.Group("/api/obat")
		{
			obat.GET("", controllers.GetAllObat)
			obat.GET("/:id", controllers.GetObatByID)
			obat.POST("", controllers.CreateObat)
			obat.PUT("/:id", controllers.UpdateObat)
			obat.DELETE("/:id", controllers.DeleteObat)
		}

		// CART
		cart := protected.Group("/cart")
		{
			cart.GET("", controllers.GetCart)
			cart.POST("", controllers.AddToCart)
			cart.PUT("/:id", controllers.UpdateCart)
			cart.DELETE("/:id", controllers.DeleteCart)
			cart.DELETE("", controllers.ClearCart)
		}

		// ORDERS
		orders := protected.Group("/orders")
		{
			orders.GET("", controllers.GetOrders)
			orders.POST("", controllers.CreateOrder)
			orders.GET("/:id", controllers.GetOrderByID)
			orders.PUT("/:id/status", controllers.UpdateOrderStatus)
		}
	}

	// REPORT
	report := r.Group("/api/admin")
	report.Use(middleware.AuthMiddleware())

	report.GET("/report", controllers.GenerateAndSaveReport)
	report.GET("/reports", controllers.GetSavedReports)
}
