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
