package config

import (
	"log"

	"github.com/joho/godotenv"
)

// Envload : check .env and load
func Envload() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(".env not found")
	}
}
