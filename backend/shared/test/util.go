package test

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func ConnectTestDB(t *testing.T) *sql.DB {
	t.Helper()

	wd, _ := os.Getwd()
	t.Logf("テストのワーキングディレクトリ: %s", wd)

	if err := godotenv.Load("../../../../../.env"); err != nil {
		log.Println("Warning: .env ファイルが見つかりませんでした")
	}

	// Get test database connection parameters from environment variables
	dbUser := getEnv("TEST_DB_USER", "test_user")
	dbPass := getEnv("TEST_DB_PASSWORD", "test_password")
	dbName := getEnv("TEST_DB_NAME", "test_db")
	dbHost := getEnv("TEST_DB_HOST", "localhost")
	dbPort := getEnv("TEST_DB_PORT", "3307")

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
