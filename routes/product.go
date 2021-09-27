package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraboy/go-fiber-api/services"
)

func ProductV1Route(c *fiber.App) {
	r := c.Group("/v1/products")
	r.Get("", services.GetAllProduct)
	r.Get("/:id", services.FindProduct)
	r.Post("", services.CreateProduct)
	r.Put("/:id", services.UpdateProduct)
	r.Delete("/:id", services.DeleteProduct)
}
