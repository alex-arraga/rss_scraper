package main

import (
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/api/middlewares"
	"github.com/alex-arraga/rss_project/internal/api/routes"
	"github.com/alex-arraga/rss_project/internal/auth"
	"github.com/alex-arraga/rss_project/internal/config"
	"github.com/alex-arraga/rss_project/internal/database/connection"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/di"
	"github.com/alex-arraga/rss_project/internal/scrapper"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq" // Import PostgresSQL driver
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Logger setup
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Starting application...")

	port, dbURL, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Configuration error")
	}

	// Db connection using pq driver
	conn, err := connection.ConnectDB(dbURL)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connet to the database")
	}
	defer conn.Close()

	db := database.New(conn)

	authService := &auth.AuthService{DB: db}
	middlewareConfig := &middlewares.MiddlewareConfig{AuthService: authService}

	go func() {
		log.Info().Msg("Starting scrapper")
		scrapper.StartScrapping(db, 10, time.Minute)
	}()

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

	c, _ := di.NewContainer(db)
	handlerConfig := handlers.NewHandlerConfig(c)

	routes.RegisterRoutes(
		router,
		*handlerConfig,
		middlewareConfig.MiddlewareAuth,
	)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Info().Msgf("Server starting on port: %s", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}
}
