package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alex-arraga/rss_project/internal/config"

	"github.com/alex-arraga/rss_project/internal/database/connection"
	"github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq" // Import PostgresSQL driver
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	port, dbURL := config.LoadConfig()

	// Db connection using pq driver
	conn := connection.ConnectDB(dbURL)

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScrapping(db, 10, time.Minute)

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

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUserByAPIKey))

	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Put("/feeds/{feedID}", apiCfg.middlewareAuth(apiCfg.handlerUpdateFeed))
	v1Router.Delete("/feeds/{feedID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeed))

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedsFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

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
