package main

import (
	"log"
	"rent-car-project/config"
	"rent-car-project/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.ConnectDB()

	r := gin.Default()

	// Register Routes
	routes.RegisterRoutes(r)

	// Route dasar (Health Check)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Car Rental API is running",
		})
	})

	// Jalankan server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
