package controllers

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"

	"github.com/GARMA-A/Projects_Go/types"
)

func UpdateTodo(c *fiber.Ctx) error {
	// Parse new todo data
	newtodo := new(types.ToDo)
	if err := c.BodyParser(newtodo); err != nil {
		return c.Status(422).JSON(fiber.Map{"error": err.Error()})
	}

	// Get ID from URL and convert to correct type (assuming it's an int)
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid ID"})
	}

	// Get collection
	collection := types.CurrentCollection(types.GlobalClient)

	// Build filter and update
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{
		"task":      newtodo.Task,
		"completed": newtodo.Completed,
	}}

	// Perform update
	result, err := collection.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to update todo"})
	}
	if result.MatchedCount == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	}

	return c.Status(200).JSON(fiber.Map{"updated_todo": newtodo})
}
