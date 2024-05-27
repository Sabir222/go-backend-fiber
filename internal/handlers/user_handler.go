package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/sabir222/go-backend-fiber/internal/models"
	"github.com/sabir222/go-backend-fiber/internal/repositories"
	"time"
)

type UserHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) HandleCreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := h.repo.CreateUser(ctx, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
