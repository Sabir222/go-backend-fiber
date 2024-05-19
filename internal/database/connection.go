package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"    // Import the PostgreSQL driver
	_ "github.com/joho/godotenv/autoload" // Import the library for loading environment variables
)

type Service interface {
	Health() map[string]string // Health method checks the health of the database
}

// service is a struct implementing the Service interface
type service struct {
	db *sql.DB // db is a pointer to the SQL database connection
}

var (
	// Define variables to hold database connection information from environment variables
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	dbInstance *service // Singleton instance of the database service
)

// New creates a new instance of the database service
func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	// Construct connection string using environment variables
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	// Establish a new database connection
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	// Store the database connection in the singleton instance
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Ping the database to check its health
	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	// Return a map indicating that the database is healthy
	return map[string]string{
		"message": "It's healthy",
	}
}
