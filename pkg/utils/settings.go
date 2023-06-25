package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	err               error
	BaseURL           string
	Environment       string
	Libp2pURL         string
	ErrEnvVarNotFound = errors.New("environment variable is not found in the .env file")
)

func InitSettings() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	BaseURL, err = getEnv("BASE_URL", "localhost:3000")
	if err != nil {
		return err
	}
	Environment, err = getEnv("ENVIRONMENT", "DEV")
	if err != nil {
		return err
	}
	Libp2pURL, err = getEnv("LIBP2P_URL", "")
	if err != nil {
		return err
	}
	return nil
}

func getEnv(envVar string, defaultValue string) (string, error) {
	value, exists := os.LookupEnv(envVar)
	if !exists {
		if defaultValue != "" {
			return "", ErrEnvVarNotFound
		} else {
			return defaultValue, nil
		}
	} else {
		return value, nil
	}
}

func IsEnvDev() bool {
	if Environment == "DEV" {
		return true
	} else {
		return false
	}
}
