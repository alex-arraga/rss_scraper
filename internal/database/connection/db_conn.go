package connection

import (
	"database/sql"
	"fmt"
)

// ConnectDB establishes a persistent connection to the PostgreSQL database using the given URL.
// The returned *sql.DB can be reused for the application's lifetime.
func ConnectDB(dbURL string) (*sql.DB, error) {
	// Attempt to open a connection pool
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Validate the connection with Ping to ensure it's reachable
	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Return the persistent connection pool
	return conn, nil
}
