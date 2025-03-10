package routes

import (
	"go-crud-app/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有路由
func SetupRoutes(r *gin.Engine) {
	// API分组
	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("/", controllers.GetAllUsers)
			users.GET("/:id", controllers.GetUser)
			users.POST("/", controllers.CreateUser)
			users.PUT("/:id", controllers.UpdateUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}
} 