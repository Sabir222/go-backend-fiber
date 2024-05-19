package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabir222/go-backend-fiber/pkg/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/user", handlers.UserHandler)
}
