package main

import (
	"backend/config"
	"backend/db"
	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config (DB, env vars, etc.)
	config.LoadConfig()

	db.InitDB()

	// Initialize Gin
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.Default())

	// Register routes
	routes.SetupRoutes(router)

	// Start server
	router.Run(":8080")
}
