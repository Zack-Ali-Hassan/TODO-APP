package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        string `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error in .env")
	}
	app := fiber.New()
	todos := []Todo{}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})
	app.Post("/api/todo", func(c *fiber.Ctx) error {
		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "please enter a body ?"})
		}
		todo.ID = fmt.Sprint(len(todos) + 1)
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)
	})
	PORT := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + PORT))
}
