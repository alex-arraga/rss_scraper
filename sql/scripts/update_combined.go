package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	outputFile := "./sql/combined_schema/combined.sql"

	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the enviroment")
	}

	// Comamand to export schemas using pg_dump
	cmd := exec.Command("pg_dump",
		"--schema-only",
		"--no-owner",
		"--file", outputFile,
		dbURL,
	)

	// Ejecuta el comando y captura errores
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error executing pg_dump: %v", err)
	}

	log.Println("File updated. Ready to use sqlc")
}
