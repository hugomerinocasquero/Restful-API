package main

import (
	"github.com/gin-gonic/gin"

	"api/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterUserRoutes(router)

	router.Run(":8080")
}
