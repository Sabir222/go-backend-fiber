package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) SetupRoutes() {
	s.App.Get("/health", s.HandleCheckHealth)
	s.App.Get("/", s.HelloWorldHandler)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) HandleCheckHealth(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
