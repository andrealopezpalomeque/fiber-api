package routes

import (
	"github.com/andrealopezpalomeque/fiber-api/database"
	"github.com/andrealopezpalomeque/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)


type User struct {
	//? This is not the model, more like a serializer
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

// CreateResponseUser takes a models.User and returns a simplified User for response.
func CreateResponseUser(user models.User) User {
	return User{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	//variable para guardar el body del request
	var user models.User

	//BODY PARSER ->  parsing the request body, which is typically in JSON format, into the user variable, which is of type models.User.
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	//crear usuario en la base de datos
	database.Database.Db.Create(&user)
	

	//respuesta simplificada del usuario
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
	
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)

	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)

}