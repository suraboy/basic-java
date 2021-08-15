package services

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"name": "Boy",
		"age":  20,
	})
}
