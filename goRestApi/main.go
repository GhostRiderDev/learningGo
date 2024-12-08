package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/goRestApi/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	SERVER_PORT := os.Getenv("SERVER_PORT")
	MONGODB_URI := os.Getenv("MONGODB_URI")

	if SERVER_PORT == "" {
		SERVER_PORT = "3000"
	}

	if MONGODB_URI == "" {
		log.Fatal("MongoDB URI not found")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGODB_URI))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	users := client.Database("goRestApi").Collection("users")

	users.InsertOne(context.TODO(), bson.D{{
		Key:   "name",
		Value: "Olvadis",
	}})

	app := fiber.New()

	app.Static("/", "public")
	app.Use(cors.New())

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "Este es un usario valido",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		var userToSave model.User

		c.BodyParser(&userToSave)

		result, err := users.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: userToSave.Name,
		}, {
			Key:   "age",
			Value: userToSave.Age,
		}})

		if err != nil {
			log.Fatal("No se pudo guardar el usuario")
		}

		log.Printf("%s\n", result)

		return c.JSON(&userToSave)
	})

	app.Listen(":" + SERVER_PORT)
	fmt.Println("Server is running on port 3000")
}
