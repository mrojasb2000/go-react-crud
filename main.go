package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mrojasb2000/go-react-crud/models"
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
		var users []models.User

		coll := client.Database(dBName).Collection("users")
		result, err := coll.Find(context.TODO(), bson.M{})

		if err != nil {
			panic(err)
		}

		for result.Next(context.TODO()) {
			var user models.User
			result.Decode(&user)
			users = append(users, user)
		}
		return c.JSON(&fiber.Map{
			"users": users,
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		c.BodyParser(&user)

		_, err := coll.InsertOne(
			context.TODO(),
			bson.D{
				{Key: "name", Value: user.Name},
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
