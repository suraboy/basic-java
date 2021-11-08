package http

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (hdl *HTTPHandler) HealthCheck(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
