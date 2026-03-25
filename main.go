package main

import (
	"log"
	"os"

	"medigo-be/config"
	"medigo-be/models"
	"medigo-be/routes"

	"github.com/gin-contrib/cors"
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

	// CORS middleware untuk Gin
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://medigo-fe.vercel.app/medicines"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
