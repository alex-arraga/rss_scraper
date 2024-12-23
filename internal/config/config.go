package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import PostgresSQL driver
)

// LoadConfig load enviroment variables and validates them.
func LoadConfig() (string, string, error) {
	// Define the .env file as a constant
	const envFile string = ".env"

	// Load enviroment variables from the .env file
	if err := godotenv.Load(envFile); err != nil {
		return "", "", fmt.Errorf("failed to load %s file: %w", envFile, err)
	}

	// Retrieval and validation of required variables
	port, err := getEnv("PORT")
	if err != nil {
		return "", "", fmt.Errorf("error reading PORT: %w", err)
	}

	dbURL, err := getEnv("DB_URL")
	if err != nil {
		return "", "", fmt.Errorf("error reading DB_URL: %w", err)
	}

	return port, dbURL, nil
}

// getEnv retrieves an environment variable or returns an error if it's missing.
func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s is not found in the environment", key)
	}
	return value, nil
}
