package main

import (
	"log"

	"TZ-GREEN-API_/internal/app"
	"TZ-GREEN-API_/internal/config"
)

func main() {
	cfg := config.GetConfig()
	App := app.App{Config: cfg}

	if err := App.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
