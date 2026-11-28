package routes

import (
	"github.com/gin-gonic/gin"

	"api/controllers"
	"api/middleware"
)

func RegisterUserRoutes(router *gin.Engine) {
	users := router.Group("/users", middleware.RequireAPIKey())
	{
		users.GET("", controllers.GetUsers)
		users.POST("", controllers.CreateUser)
		users.GET("/:id", controllers.GetUserByID)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}
}
