// Request handlers (controllers) that handle incoming HTTP requests and send responses back to the client.
// Handles requests related to user operations, such as registration, login, and profile management.
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sabir222/go-backend-fiber/pkg/models"
)

type User struct {
	user models.User
}

func AddUserHandler(c *fiber.Ctx) {
	newUser :=
}
