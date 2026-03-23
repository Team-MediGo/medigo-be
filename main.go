package main

import (
	"log"
	"medigo-be/config"
	"medigo-be/models"
	"medigo-be/routes"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan")
	}
	mux := http.NewServeMux()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://be.mdgo.my.id"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)

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
	log.Fatal(http.ListenAndServe(":8080", handler))
}
