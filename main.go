package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pobyzaarif/belajar-go-cli/config"
)

func main() {
	// Load .env dari root project
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set")
	}

	// Init database
	db := config.InitDatabase(dsn)
	defer db.Close()

}
