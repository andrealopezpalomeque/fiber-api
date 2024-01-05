package routes

import (
	"errors"

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
	
	var user models.User //variable para guardar el body del request

	//BODY PARSER ->  parsing the request body, which is typically in JSON format, into the user variable, which is of type models.User.
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	database.Database.Db.Create(&user) //crear usuario en la base de datos
	
	responseUser := CreateResponseUser(user) //respuesta simplificada del usuario

	return c.Status(200).JSON(responseUser)
	
}

func GetUsers(c *fiber.Ctx) error {
	
	users := []models.User{} //slice para guardar los usuarios de la base de datos

	database.Database.Db.Find(&users) //busco los usuarios en la base de datos
	
	responseUsers := []User{} //slice para guardar la respuesta simplificada de los usuarios

	for _, user := range users {
		
		responseUser := CreateResponseUser(user) //creo una respuesta simplificada por CADA usuario
		
		responseUsers = append(responseUsers, responseUser) //agrego la respuesta simplificada al slice de respuesta
	}

	return c.Status(200).JSON(responseUsers)

}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}


func GetUser(c *fiber.Ctx) error {
	
	id, err := c.ParamsInt("id") //recupero el id de params

	var user models.User //un usuario vacio

	if err != nil {
		return c.Status(400).JSON("Please ensure id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)


}