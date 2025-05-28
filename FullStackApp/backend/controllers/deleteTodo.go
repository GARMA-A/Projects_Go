package controllers

import (
	"fmt"
	"slices"

	"github.com/gofiber/fiber/v2"

	"github.com/GARMA-A/Projects_Go/types"
)

func DeleteTodo(c *fiber.Ctx) error {
	todos := types.Todos
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			types.Todos = slices.Delete(types.Todos, i, i+1)
			return c.Status(200).JSON(fiber.Map{"message": "todo deleted successfully"})
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
}
