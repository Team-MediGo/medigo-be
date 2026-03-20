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
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, pakai environment variable system")
	}

	// Koneksi database
	config.ConnectDB()

	// Auto migrate
	config.DB.AutoMigrate(&models.Obat{})

	// Setup router
	r := gin.Default()
	routes.SetupRoutes(r)

	// Jalankan server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
