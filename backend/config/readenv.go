package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadENV() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
