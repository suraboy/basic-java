package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraboy/go-fiber-api/pkg/logger"
)

type Middleware interface {
	LogRequest(c *fiber.Ctx) error
}

type middleware struct{}

func NewMiddleware() Middleware {
	return &middleware{}
}

func (mdw *middleware) LogRequest(c *fiber.Ctx) error {
	IP := c.IP()
	method := c.Method()
	path := c.Path()
	body := string(c.Request().Body())
	query := string(c.Request().URI().QueryString())
	logger.AppLogger.Request(IP, method, path, body, query)
	return c.Next()
}
