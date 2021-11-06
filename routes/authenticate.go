package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraboy/go-fiber-api/services"
	jwtware "github.com/gofiber/jwt/v3"
)

func AuthenticateV1Route(c *fiber.App) {
	// JWT Middleware
	c.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	// Restricted Routes
	r := c.Group("/v1/oauth")
	r.Get("/login", services.Login)
}
