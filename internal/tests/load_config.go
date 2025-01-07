package tests

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadTestConfig() (string, error) {
	const envFile string = "C:/Users/arrag/Desktop/Dev/Go/rss_project/internal/tests/.env.test"

	if err := godotenv.Load(envFile); err != nil {
		return "", fmt.Errorf("failed to load %s file: %w", envFile, err)
	}

	db, err := getEnv("DB_URL_TEST")
	if err != nil {
		log.Fatal("error loading env:", err)
	}

	return db, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s is not found in the environment", key)
	}
	return value, nil
}
