package controllers

import (
	"github.com/GARMA-A/Projects_Go/types"
	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	todo := new(types.ToDo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.ID == 0 {
		todo.ID = len(types.Todos) + 1
	}
	if todo.Task == "" {
		return c.Status(400).JSON(fiber.Map{"error": "task body is required"})
	}

	// Add the todo to the database (mocked here)
	types.Todos = append(types.Todos, *todo)

	return c.Status(fiber.StatusCreated).JSON(todo)
}
