package v1

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"name": "Grame",
		"age":  20,
	})
}
