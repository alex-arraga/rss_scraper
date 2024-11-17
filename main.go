package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello RSS")

	// Export the variables of .env in the project
	godotenv.Load(".env")

	// Read env "PORT"
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found")
	}

	// Create router and server
	router := chi.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port: %s", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
