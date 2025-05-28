package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/GARMA-A/Projects_Go/types"
)

func CreateTodo(c *fiber.Ctx) error {
	todo := new(types.ToDo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.ID == "" {
		todo.ID = uuid.New().String() // Generate a new ID if not provided
	}
	if todo.Task == "" {
		return c.Status(400).JSON(fiber.Map{"error": "task body is required"})
	}

	collection := types.CurrentCollection(types.GlobalClient)
	_, err := collection.InsertOne(c.Context(), todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(todo)
}
