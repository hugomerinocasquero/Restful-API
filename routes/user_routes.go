package routes

import (
	"github.com/gin-gonic/gin"

	"api/controllers"
)

func RegisterUserRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.CreateUser)
	router.GET("/users/:id", controllers.GetUserByID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)
}
