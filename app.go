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
	server.Use(middleware.LoggerMiddleware())

	routes.RegisterRoutes(server)
	routes.RegisterDocsRoutes(server)

	server.Run(":8080")
}
