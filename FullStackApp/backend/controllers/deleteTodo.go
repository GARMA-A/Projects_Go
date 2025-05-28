package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/GARMA-A/Projects_Go/types"
)

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	collection := types.CurrentCollection(types.GlobalClient)
	_, err := collection.DeleteOne(c.Context(), fiber.Map{"id": id})
	if err != nil {
		fmt.Println("Error deleting todo:", err)
		return c.Status(500).JSON(fiber.Map{"error": "failed to delete todo"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "todo deleted successfully"})
}
