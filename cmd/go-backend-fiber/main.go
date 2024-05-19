package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "go backend v1",
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})
	app.Static("/", "./public/index.html")
	app.Get("/error", func(c *fiber.Ctx) error {
		return fiber.NewError(398, "Tf even is that")
	})
	app.Listen(":3000")
}
