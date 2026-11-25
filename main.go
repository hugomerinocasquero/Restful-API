package main

import (
	"github.com/gin-gonic/gin"

	"api/database"
	"api/routes"
)

func main() {
	database.Connect()

	router := gin.Default()

	routes.RegisterUserRoutes(router)

	router.Run(":8080")
}
