package main

import (
	"log"
	"os"

	"gin-simple.com/config"
	"gin-simple.com/handlers"
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

	userHandler := handlers.NewUserHandler(db)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userHandler.GetAll)
		v1.GET("/users/:id", userHandler.GetByID)
		v1.POST("/users", userHandler.Create)
		v1.PUT("/users/:id", userHandler.Update)
		v1.DELETE("/users/:id", userHandler.Delete)
	}

	r.Run(":" + os.Getenv("PORT"))
}
