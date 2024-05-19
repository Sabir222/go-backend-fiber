# Go Backend using Fiber PostgreSQL JWT (In progress)

## Project Structure

This project is organized as follows:

```plaintext
myproject/
├── cmd/
│   └── main.go
├── pkg/
│   ├── handlers/
│   │   └── user_handler.go
│   ├── models/
│   │   └── user.go
│   ├── repositories/
│   │   └── user_repository.go
│   ├── services/
│   │   └── user_service.go
│   ├── middlewares/
│   │   └── auth_middleware.go
│   └── utils/
│       └── response.go
├── internal/
│   ├── app/
│   │   └── router.go
│   ├── database/
│   │   └── connection.go
│   └── auth/
│       └── jwt.go
├── config/
│   └── config.go
├── scripts/
│   └── migrate.sh
├── docs/
│   └── api.md
├── Makefile
├── go.mod
└── go.sum

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

test
