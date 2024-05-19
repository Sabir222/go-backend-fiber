package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabir222/go-backend-fiber/internal/server"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "go backend v1",
	})

	server.SetupRoutes(app)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(":3000")
}
