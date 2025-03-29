package main

import (
	"log"
	"movie/internal/config"
	"movie/internal/routes"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}

	db, err := config.NewDB(cfg)
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
	}

	router := routes.SetupRouter(db)
	if err := router.Run("8080"); err != nil {
		log.Printf("Error starting server: %v", err)
	}
}
