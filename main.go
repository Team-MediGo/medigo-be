package main

import (
	"log"
	"medigo-be/config"
	"medigo-be/models"
	"medigo-be/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan")
	}

	config.ConnectDB()
	config.ConnectCloudinary()
	config.DB.AutoMigrate(&models.Obat{})

	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
