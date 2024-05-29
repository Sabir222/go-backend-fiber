package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string // Health method checks the health of the database
	GetDb() *sql.DB
}

// service is a struct implementing the Service interface
type service struct {
	db *sql.DB // db is a pointer to the SQL database connection
}

var (
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

	// create user table

	dbInstance.CreateUserTable()

	return dbInstance
}

// createUserTable

func (s *service) CreateUserTable() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if err != nil {

		log.Fatalf(fmt.Sprintf("cannot create extension: %v", err))
	}

	createUserTableQuery := `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		full_name VARCHAR(50),
		password_hash VARCHAR(255) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		update_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		role VARCHAR(10) DEFAULT 'user' CHECK (role IN ('admin', 'user'))
	)`
	_, err = s.db.ExecContext(ctx, createUserTableQuery)
	if err != nil {
		log.Fatalf(fmt.Sprintf("cannot create table: %v", err))
	}
}

func (s *service) GetDb() *sql.DB {
	return s.db
}

// Using context checking  the health of the database connection
// use use this in a route to check the health of the database
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
