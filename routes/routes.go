package routes

import (
	"database/sql"

	"gin-simple.com/controllers"
	"gin-simple.com/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", userController.Create)
			users.GET("/:id", userController.GetByID)
		}
	}
}
