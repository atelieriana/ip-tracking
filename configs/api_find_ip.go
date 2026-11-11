package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ApiEndPoint() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("API_ENDPOINT")
}

func ApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("API_KEY")
}
