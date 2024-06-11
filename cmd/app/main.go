package main

import (
	"back-user-api/api/router"
	"back-user-api/config"
	"back-user-api/internal/database"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	fmt.Println("START")
	// Load DB configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	fmt.Println(cfg.Port)
	fmt.Println(cfg.DatabaseURL)

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error during conecting to database: %v", err)
	}
	defer db.Close()

	r := router.SetupRouter(db)

	log.Printf("Server is running on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
