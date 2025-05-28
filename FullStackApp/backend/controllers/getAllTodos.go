package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/GARMA-A/Projects_Go/types"
)

func GetAllTodos(c *fiber.Ctx) error {
	collection := types.CurrentCollection(types.GlobalClient)
	var todos []types.ToDo

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve todos",
		})
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo types.ToDo
		if err := cursor.Decode(&todo); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to decode todo",
			})
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}
