package main

import (
	"log"

	"github.com/andrealopezpalomeque/fiber-api/database"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API!")
}

func main() {

	database.ConnectDb() //connect to the database

	app := fiber.New()

	app.Get("/api", welcome)


	log.Fatal(app.Listen(":3000"))
}