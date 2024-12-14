package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/api"
	"github.com/alex-arraga/rss_project/internal/config"
	"github.com/alex-arraga/rss_project/internal/database/connection"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/scrapper"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq" // Import PostgresSQL driver
)

func main() {
	port, dbURL := config.LoadConfig()

	// Db connection using pq driver
	conn := connection.ConnectDB(dbURL)

	db := database.New(conn)
	apiCfg := api.APIConfig{
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

	v1Router := chi.NewRouter()
	router.Mount("/v1", v1Router)

	v1Router.Get("/healthz", apiCfg.HandlerReadiness)
	v1Router.Get("/err", apiCfg.HandlerErr)
	v1Router.Post("/users", apiCfg.HandlerCreateUser) //createUser
	v1Router.Get("/users", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUserByAPIKey))

	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)
	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Put("/feeds/{feedID}", apiCfg.MiddlewareAuth(apiCfg.HandlerUpdateFeed))
	v1Router.Delete("/feeds/{feedID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeed))

	v1Router.Post("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedsFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollows))

	v1Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostsForUser))

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
