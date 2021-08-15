package routes

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoute(c *fiber.App) {
	r := c.Group("/v1/users")
	r.Get("", func(c *fiber.Ctx) error {
		return c.SendString("User V2")
	})

	//e.POST("/v1/login",api.LoginUser)
}
