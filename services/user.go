package services

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	postgreSQL "github.com/suraboy/go-fiber-api/config"
	"github.com/suraboy/go-fiber-api/models"
)

func GetAllUser(c *fiber.Ctx) error {
	postgreSQL := postgreSQL.PostgresConnection()
	var users []models.Users
	postgreSQL.Find(&users)
	for _, user := range users{
		fmt.Println(user.Email, user.Username, user.Password)
	}
	return c.JSON(fiber.Map{"data": users})
}
