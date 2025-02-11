package main

import (
	"log"
	"os"

	"gin-simple.com/config"
	"gin-simple.com/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// init is executed first to load env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Initialize DB
	db := config.InitDB()
	defer db.Close()

	// Initialize Gin
	r := gin.Default()

	// Setup Routes
	routes.SetupRoutes(r, db)

	r.Run(":" + os.Getenv("PORT"))
}
