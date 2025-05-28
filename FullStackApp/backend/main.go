package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/GARMA-A/Projects_Go/controllers"
	"github.com/GARMA-A/Projects_Go/types"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	Port := os.Getenv("PORT")

	client, err := types.ConnectToMongoDB()
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	} else {
		log.Println("Connected to MongoDB successfully")
	}
	defer client.Disconnect(context.Background())
	types.GlobalClient = client

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "Hello, World!"})
	})

	app.Get("/todos", controllers.GetAllTodos)

	app.Post("/create_todo", controllers.CreateTodo)

	app.Put("/update_todo/:id", controllers.UpdateTodo)

	app.Delete("/delete_todo/:id", controllers.DeleteTodo)

	log.Fatal(app.Listen(":" + Port))
}
