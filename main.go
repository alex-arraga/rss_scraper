package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/alex-arraga/rss_project/internal/api/routes"
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
	port, dbURL := config.LoadConfig()

	// Db connection using pq driver
	conn := connection.ConnectDB(dbURL)

	db := database.New(conn)

	// ! Temp
	apiCfg := api.APIConfig{
		DB: db,
	}

	serviceConfig := services.ServicesConfig{
		DB: db,
	}

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

	routes.RegisterRoutes(router, &apiCfg, &serviceConfig)

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
