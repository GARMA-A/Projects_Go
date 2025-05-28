package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/GARMA-A/Projects_Go/types"
)

func UpdateTodo(c *fiber.Ctx) error {
	todos := types.Todos
	newtodo := new(types.ToDo)
	if err := c.BodyParser(newtodo); err != nil {
		return c.Status(422).JSON(fiber.Map{"error": err})
	}
	id := c.Params("id")

	for i, todo := range todos {
		if fmt.Sprint(todo.ID) == id {
			todos[i].Task = newtodo.Task
			todos[i].Completed = newtodo.Completed

			return c.Status(200).JSON(fiber.Map{"updated_todo": newtodo})
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
}
