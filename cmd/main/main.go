package main

import (
	"log"
	"os"
	"path/filepath"
	"test_orders/internal/app/get"
	"test_orders/internal/config"
	"test_orders/internal/storage/postgres"

	"github.com/joho/godotenv"
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fullPath := filepath.Join(path, ".env")
	err = godotenv.Load(fullPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cfg := config.MustLoad()
	db, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	get.GetOrders(db)
}
