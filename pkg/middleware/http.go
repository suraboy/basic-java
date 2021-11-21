package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	responseBody "go-fiber-api/pkg/util"
)

type Middleware interface {
	JWTMiddleware(*fiber.Router) fiber.Router
}

type middleware struct{}

func NewMiddleware() *middleware {
	return &middleware{}
}

func (mdw *middleware) JWTMiddleware(router fiber.Router) fiber.Router {
	return router.Use(jwtware.New(jwtware.Config{
		SuccessHandler: AuthSuccess,
		ErrorHandler:   AuthError,
		SigningKey:     []byte("secret"),
		SigningMethod:  "HS256",
	}))
}

func AuthError(c *fiber.Ctx, e error) error {
	return responseBody.SendUnauthorized()
}
func AuthSuccess(c *fiber.Ctx) error {
	return c.Next()
}
