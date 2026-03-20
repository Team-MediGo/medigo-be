package handler

import (
	"medigo-be/config"
	"medigo-be/models"
	"medigo-be/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var app *gin.Engine

func init() {
	godotenv.Load()
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Obat{})

	app = gin.Default()
	routes.SetupRoutes(app)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
