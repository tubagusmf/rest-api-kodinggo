package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadWithGodotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error reading config file .env, %s", err)
	}
}
