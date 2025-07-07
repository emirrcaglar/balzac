// cmd/server/main.go
package main

import (
	"balzac/config"
	"balzac/routes"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	cfg := config.Load()
	// Validate required config
	if cfg.GoogleClientID == "" || cfg.GoogleClientSecret == "" {
		log.Fatal("GOOGLE_CLIENT_ID and GOOGLE_CLIENT_SECRET must be set")
	}
	if cfg.SessionSecret == "" {
		log.Fatal("SESSION_SECRET must be set")
	}
	// Setup routes
	router := routes.SetupRoutes(cfg)
	// Start server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	log.Printf("Frontend URL: %s", cfg.FrontendURL)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
