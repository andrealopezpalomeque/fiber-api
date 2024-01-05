package main

import (
	"log"

	"github.com/andrealopezpalomeque/fiber-api/database"
	"github.com/andrealopezpalomeque/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API!")
}

func setupRoutes(app *fiber.App) {
	//welcome endpoint
	app.Get("/api", welcome)

	//USER ENDPOINTS
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
}

func main() {

	database.ConnectDb() //connect to the database

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}