package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import PostgresSQL driver
)

func LoadConfig() (string, string) {
	// Export the variables of .env in the project
	godotenv.Load(".env")

	// Read env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the enviroment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the enviroment")
	}

	return port, dbURL
}
