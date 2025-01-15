package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq" // Import PostgresSQL driver
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/alex-arraga/rss_project/internal/api/handlers"
	"github.com/alex-arraga/rss_project/internal/api/middlewares"
	"github.com/alex-arraga/rss_project/internal/api/routes"
	"github.com/alex-arraga/rss_project/internal/config"
	"github.com/alex-arraga/rss_project/internal/database/connection"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/di"
	"github.com/alex-arraga/rss_project/internal/logger"
	"github.com/alex-arraga/rss_project/internal/scrapper"
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

	// Container that inject dependencies (db -> services)
	container := di.NewContainer(db)

	middlewareConfig := &middlewares.MiddlewareConfig{
		AuthService: container.AuthSerive,
	}

	go func() {
		log.Info().Msg("Starting scrapper")
		scrapper.StartScrapping(db, 10, time.Minute)
	}()

	go func() {
		start := time.Now()
		logger.RecordHTTPRequests(time.Since(start))

		// Initialize prometheus server
		err = logger.StartPrometheus()
		if err != nil {
			log.Fatal().Msgf("Prometheus error: %v", err)
		}
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

	handlerConfig := handlers.NewHandlerConfig(container)

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

	log.Info().Msgf("Application server starting on port: %s", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Application server failed")
	}
}
