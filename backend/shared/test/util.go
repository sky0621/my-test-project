package test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func ConnectTestDB(t *testing.T) *sql.DB {
	t.Helper()

	dbHost := os.Getenv("TEST_DB_HOST")
	dbPort := os.Getenv("TEST_DB_PORT")
	dbName := os.Getenv("TEST_DB_NAME")
	dbUser := os.Getenv("TEST_DB_USER")
	dbPass := os.Getenv("TEST_DB_PASSWORD")

	// Create DSN for MySQL connection
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPass, dbHost, dbPort, dbName,
	)

	// Open database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	return db
}

// getEnv returns the value of the environment variable or a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
