package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabir222/go-backend-fiber/internal/handlers"
	"github.com/sabir222/go-backend-fiber/internal/repositories"
)

func (s *FiberServer) SetupRoutes() {

	userRepository := repositories.NewUserRepository(s.db.GetDb())
	userHandler := handlers.NewUserHandler(userRepository)
	s.App.Get("/health", s.HandleCheckHealth)
	s.App.Post("/user", userHandler.HandleCreateUser)
}

func (s *FiberServer) HandleCheckHealth(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
