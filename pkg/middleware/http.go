package middleware

import (
	"go-fiber-api/config"
	"go-fiber-api/pkg/logger"
	"go-fiber-api/pkg/util"

	"github.com/gofiber/fiber/v2"
)

type Middleware interface {
	Logger() fiber.Handler
	ApiKeyAccess(c *fiber.Ctx) error
}

type middleware struct{}

func NewMiddleware() Middleware {
	return &middleware{}
}

func (mdw *middleware) Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		logger.HTTP(logger.HTTPLog{
			IP:           c.IP(),
			Method:       c.Method(),
			Path:         c.Path(),
			RequestBody:  string(c.Request().Body()),
			Query:        string(c.Request().URI().QueryString()),
			ResponseBody: string(c.Response().Body()),
			Status:       c.Response().StatusCode(),
		})
		return err
	}

}

func (mdw *middleware) ApiKeyAccess(c *fiber.Ctx) error {
	header := c.Get("x-api-key")
	apiKey := config.GetViper().App.APIKey

	if header != apiKey {
		return c.Status(401).JSON(util.CustomResponse(401, "Unauthorized"))
	}

	return c.Next()
}
