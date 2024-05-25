# Go Backend using Fiber PostgreSQL JWT (In progress)

## Project Structure

This project is organized as follows:

```plaintext
myproject/
Add project tree here later.
```

## Connect go to database

- Go to docker hub and sign-up : [DockerHub](https://hub.docker.com/_/postgres).
- Login to docker-hub in your terminal by using `docker login` then enter username and pw or using [token](https://hub.docker.com/settings/security) docker login -u (username)` then enter your token.

To run a PostgreSQL container use the following command:

```bash
docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=MySecretPassword -d postgres
```

- `--name` to name it.
- `-p 5432:5432` Expose container's port 5432 to host's port 5432.
- `-e POSTGRES_PASSWORD=MySecretPassword ` Set the environment variable `POSTGRES_PASSWORD` to "MySecretPassword".
- `-d` Run the container in detached mode (`-d` flag).
- `postgres` what image we want to pull.

```
# List all Docker containers.
docker ps -a

# Start a Docker container named <docker container name>.
docker start <docker container name>

# Open a terminal within the Docker container named (express-typescript-postgres-docker).
sudo docker exec -it express-typescript-postgres-docker bash

# Access the PostgreSQL database.
psql -U postgres

# List all databases.
\l

# Connect to a specific database like 'backenddb'.
\c backenddb

```

## Link Project with Database

Install packages

```bash

go get github.com/joho/godotenv

 go get github.com/jackc/pgx/v5

```

Create a `.env` file in the root directory:

```plaintext

PORT=8080
APP_ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_DATABASE=go_fiber
DB_USERNAME=postgres
DB_PASSWORD=DB_PASSWORD

```
- Add connection to database in `internal/database/connection.go`:

```go

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

	dbInstance.CreateUserTable()
	return dbInstance
}


// createUserTable

func (s *service) CreateUserTable() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := s.db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		email VARCHAR(50)
	)`)
	if err != nil {
		log.Fatalf(fmt.Sprintf("cannot create table: %v", err))
	}
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
```

- Create Fiber server in `internal/server/server.go`:

```go
package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabir222/go-backend-fiber/internal/database"
)


type FiberServer struct {
	*fiber.App
	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "Fiber",
			AppName:      "go-backend-fiber",
		}),
		db: database.New(),
	}

	return server
}
```

- Add the following code to `cmd/main.go`:

```go
package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sabir222/go-backend-fiber/internal/server"
	"os"
	"strconv"
)

func main() {

	server := server.New()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	fmt.Println(port)
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
```
