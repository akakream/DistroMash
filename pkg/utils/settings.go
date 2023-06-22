package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BaseURL string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	BaseURL = getEnv("BASE_URL", "localhost:3000")
}

func getEnv(envVar string, defaultValue string) string {
	value, exists := os.LookupEnv(envVar)
	if !exists {
		return defaultValue
	} else {
		return value
	}
}
