package main

import (
	"log"
	"movie/internal/config"
	"movie/internal/routes"
)

// Gitlab is shit

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := config.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	address := ":8080"
	router := routes.SetupRouter(db)
	log.Printf("Server starting on %s", address)
	if err := router.Run(address); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
