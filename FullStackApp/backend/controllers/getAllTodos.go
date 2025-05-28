package controllers

import (
	"github.com/GARMA-A/Projects_Go/types"
	"github.com/gofiber/fiber/v2"
)

func GetAllTodos(c *fiber.Ctx) error {
	todos := fiber.Map{"todos": types.Todos}
	return c.Status(fiber.StatusOK).JSON(todos)
}
