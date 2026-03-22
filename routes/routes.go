package routes

import (
	"medigo-be/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	obat := r.Group("/api/obat")
	{
		obat.GET("", controllers.GetAllObat)
		obat.GET("/:id", controllers.GetObatByID)
		obat.POST("", controllers.CreateObat)
		obat.PUT("/:id", controllers.UpdateObat)
		obat.DELETE("/:id", controllers.DeleteObat)
	}
}
