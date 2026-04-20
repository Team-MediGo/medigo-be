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
		// Obat
		obat := protected.Group("/api/obat")
		{
			obat.GET("", controllers.GetAllObat)
			obat.GET("/:id", controllers.GetObatByID)
			obat.POST("", controllers.CreateObat)
			obat.PUT("/:id", controllers.UpdateObat)
			obat.DELETE("/:id", controllers.DeleteObat)
		}

		// Cart
		cart := protected.Group("/cart")
		{
			cart.GET("", controllers.GetCart)
			cart.POST("", controllers.AddToCart)
			cart.PUT("/:id", controllers.UpdateCart)
			cart.DELETE("/:id", controllers.DeleteCart)
			cart.DELETE("", controllers.ClearCart)
		}

		// Orders
		orders := protected.Group("/orders")
		{
			orders.GET("", controllers.GetOrders)
			orders.POST("", controllers.CreateOrder)
			orders.GET("/:id", controllers.GetOrderByID)
			orders.PUT("/:id/status", controllers.UpdateOrderStatus)
		}
	}

		// 🔥 ADMIN ROUTES
		admin := r.Group("/api/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
		admin.GET("/report", controllers.GenerateAndSaveReport) // generate + simpan
		admin.GET("/reports", controllers.GetSavedReports)      // ambil semua
		}