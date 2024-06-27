package main

import (
	"log"

	"TZ-GREEN-API_/internal/app"
	"TZ-GREEN-API_/internal/config"
)

func main() {
	cfg := config.GetConfig()
	//db := db.DataBase{}
	//db.GetStorage(cfg)
	//migration.Migrations(cfg, db.DB)
	App := app.App{Config: cfg}
	//_ := App.Start()

	if err := App.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
