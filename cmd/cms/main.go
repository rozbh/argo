package main

import (
	"log"
	"net/http"

	"github.com/rozbh/argo/internal/config"
	"github.com/rozbh/argo/internal/routes"
)

func main() {
	// Load configuration (e.g., port number)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize routes
	router := routes.InitRoutes()

	// Start the server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
