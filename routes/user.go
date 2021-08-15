package routes

import (
	"github.com/gofiber/fiber/v2"
	service "github.com/suraboy/go-fiber-api/services/v1"
)

func UserV1Route(c *fiber.App) {
	r := c.Group("/v1/users")
	r.Get("", service.GetAllUser)
	//e.POST("/v1/login",api.LoginUser)
}
