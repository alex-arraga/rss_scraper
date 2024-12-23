package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/api/middlewares"
	"github.com/alex-arraga/rss_project/internal/api/routes"
	"github.com/alex-arraga/rss_project/internal/auth"
	"github.com/alex-arraga/rss_project/internal/config"
	"github.com/alex-arraga/rss_project/internal/database/connection"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/scrapper"
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq" // Import PostgresSQL driver
)

func main() {
	port, dbURL, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Configuration error: %w", err)
	}

	// Db connection using pq driver
	conn := connection.ConnectDB(dbURL)

	db := database.New(conn)

	authService := &auth.AuthService{DB: db}
	middlewareConfig := &middlewares.MiddlewareConfig{AuthService: authService}

	go scrapper.StartScrapping(db, 10, time.Minute)

	// Create router and server
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	routes.RegisterRoutes(
		router,
		&services.ServicesConfig{DB: db},
		middlewareConfig.MiddlewareAuth,
	)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port: %s", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
