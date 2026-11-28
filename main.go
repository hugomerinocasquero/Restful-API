package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"api/database"
	"api/middleware"
	"api/routes"
)

func main() {
	database.Connect()

	router := gin.Default()

	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal("could not set trusted proxies: ", err)
	}

	limiter := middleware.NewRateLimiter(5, 10)
	router.Use(limiter.Middleware())
	router.Use(middleware.LimitBodySize())

	routes.RegisterUserRoutes(router)

	router.Run(":8080")
}
