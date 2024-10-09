package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Server started....")
	err := godotenv.Load(".env")
	if os.Getenv("ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading environment: ", err)
		}
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOption := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to mongodb success....")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:5173/",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	app.Get("/api/todos", getTodos)
	app.Get("/api/todo/:id", getTodo)
	app.Post("/api/todo", createTodo)
	app.Patch("/api/todo/:id", updateTodo)
	app.Delete("/api/todo/:id", deleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}
	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}
	log.Fatal(app.Listen(":" + port))

}

func getTodos(c *fiber.Ctx) error {
	todos := []Todo{}

	result, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		todo := Todo{}
		if err := result.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.Status(200).JSON(todos)

}

func getTodo(c *fiber.Ctx) error {
	var todo Todo
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Invalid todo id!..."})
	}
	filter := bson.M{"_id": objectId}
	err = collection.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(todo)

}
func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Please enter a body of todo..."})
	}
	inserData, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	todo.ID = inserData.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Invalid todo id!..."})
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"msg": "Updated successfully"})
}
func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Invalid todo id!..."})
	}
	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"msg": "Deleted successfully"})
}
