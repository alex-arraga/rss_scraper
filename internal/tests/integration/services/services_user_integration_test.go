package integration_services

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/alex-arraga/rss_project/internal/database/connection"
	database "github.com/alex-arraga/rss_project/internal/database/sqlc"
	"github.com/alex-arraga/rss_project/internal/services"
	"github.com/alex-arraga/rss_project/internal/tests"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestDB struct {
	db      *database.Queries
	conn    *sql.DB
	cleanup func(t *testing.T, db *sql.DB)
}

// Función helper para limpiar la base de datos
func cleanupDatabase(t *testing.T, db *sql.DB) {
	t.Helper()
	_, err := db.Exec("TRUNCATE TABLE users CASCADE")
	require.NoError(t, err, "Failed to cleanup database")
}

// setupTestDB ahora retorna una estructura TestDB con funciones de limpieza
func setupTestDB(t *testing.T) *TestDB {
	t.Helper()

	// Cargar configuración de prueba
	dbURL, err := tests.LoadTestConfig()
	require.NoError(t, err, "Failed to load test config")

	// Conectar a la base de datos
	conn, err := connection.ConnectDB(dbURL)
	require.NoError(t, err, "Failed to connect to database")

	db := database.New(conn)

	return &TestDB{
		db:      db,
		conn:    conn,
		cleanup: cleanupDatabase,
	}
}

func TestIntegration_CreateUser(t *testing.T) {
	// Configurar la base de datos de prueba
	testDB := setupTestDB(t)
	defer testDB.cleanup(t, testDB.conn) // Asegurar la limpieza al finalizar

	userService := &services.UserService{DB: testDB.db}

	tests := []struct {
		testName string
		userName string
		wantErr  bool
	}{
		{
			testName: "successfully creates user",
			userName: "John Doe",
			wantErr:  false,
		},
		{
			testName: "empty username",
			userName: "",
			wantErr:  true,
		},
		{
			testName: "empty apikey",
			wantErr:  true,
		},
		{
			testName: "empty userID",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Comenzar una transacción
			tx, err := testDB.conn.BeginTx(ctx, nil)
			require.NoError(t, err)
			defer tx.Rollback() // Asegurar rollback en caso de error

			// Ejecutar la prueba
			user, err := userService.CreateUser(ctx, tt.userName)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			// Validaciones
			assert.NoError(t, err)
			assert.NotEmpty(t, user.ID)
			assert.Equal(t, tt.userName, user.Name)
			assert.NotZero(t, user.CreatedAt)

			// Verificar que el usuario existe en la base de datos
			dbUser, err := userService.GetUserByAPIKey(ctx, user.APIKey)
			assert.NoError(t, err)
			assert.Equal(t, user.ID, dbUser.ID)
			assert.Equal(t, user.APIKey, dbUser.APIKey)
			assert.Equal(t, user.Name, dbUser.Name)

			// Commit de la transacción si todo está bien
			err = tx.Commit()
			assert.NoError(t, err)
		})
	}
}
