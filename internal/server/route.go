package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) SetupRoutes() {
	s.App.Get("/health", s.HandleCheckHealth)
}

func (s *FiberServer) HandleCheckHealth(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
