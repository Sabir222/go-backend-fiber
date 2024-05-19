// Request handlers (controllers) that handle incoming HTTP requests and send responses back to the client.
// Handles requests related to user operations, such as registration, login, and profile management.
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func UserHandler(c *fiber.Ctx) error {
	return c.SendString("<div>Hello User </div>")
}
