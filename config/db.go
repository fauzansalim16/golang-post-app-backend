package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() error {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	// Read database configuration from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	// Open database connection
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return fmt.Errorf("error opening database connection: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	DB = db
	return nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
