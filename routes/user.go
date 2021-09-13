package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraboy/go-fiber-api/services"
)

func UserV1Route(c *fiber.App) {
	r := c.Group("/v1/users")
	r.Get("", services.GetAllUser)
	r.Post("", services.CreateNewProduct)
}
