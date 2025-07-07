package main

import (
	"balzac/config"
	"balzac/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config (DB, env vars, etc.)
	config.LoadConfig()

	// Initialize Gin
	router := gin.Default()

	// Register routes
	routes.SetupRoutes(router)

	// Start server
	router.Run(":8080")
}
