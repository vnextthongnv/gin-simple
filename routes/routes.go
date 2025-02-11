package routes

import (
	"database/sql"

	"gin-simple.com/controllers"
	"gin-simple.com/repositories"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configures routes for the User model.
func SetupUserRoutes(api *gin.RouterGroup, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	users := api.Group("/users")
	{
		users.POST("/", userController.Create)
		users.GET("/:id", userController.GetByID)
		users.GET("/", userController.GetAll)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}
}

// SetupRoutes sets up all routes in the application.
func SetupRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api/v1")
	{
		SetupUserRoutes(api, db)
	}
}
