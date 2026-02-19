package main

import (
	"rest-api/db"
	"rest-api/middleware"
	"rest-api/routes"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()

	db.InitDB()

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	server.Use(middleware.LoggerMiddleware())

	routes.RegisterRoutes(server)
	routes.RegisterDocsRoutes(server)

	// Handle ALL preflight OPTIONS requests so CORS headers are returned
	// instead of 404 (Gin skips middleware for unregistered paths otherwise)
	server.OPTIONS("/*any", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	server.Run(":8080")
}
