package connection

import (
	"database/sql"
	"log"
)

func connectDB(dbURL string) *sql.DB {
	// Db connection using pq driver
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}
	defer conn.Close()

	return conn
}
