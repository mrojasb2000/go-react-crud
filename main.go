package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dBName = "react-crud"
const mongoDbUrl = "mongodb://localhost:27017/" + dBName

func main() {

	app := fiber.New()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoDbUrl))
	if err != nil {
		panic(err)
	}

	coll := client.Database(dBName).Collection("users")

	// Initialize default config
	app.Use(cors.New())

	app.Static("/", "./client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"message": "users list from backend",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		_, err := coll.InsertOne(
			context.TODO(),
			bson.D{
				{Key: "name", Value: "Mavro"},
			},
		)
		if err != nil {
			panic(err)
		}
		return c.Status(201).JSON(&fiber.Map{
			"message": "User created",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server is running on port :3000")
}
