package http

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/config"
	"go-fiber-api/pkg/util"
)

func (hdl *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.JSON(util.CustomResponse(200, "Go-fiber-api 1 is OK 👋!"))
}

type ResponseVersion struct {
	Version  string `json:"version,omitempty"`
	Database string `json:"database,omitempty"`
}

func (hdl *Handler) CheckVersion(c *fiber.Ctx) error {
	var response ResponseVersion
	response.Version = config.GetViper().App.AppVersion
	response.Database = "OK"
	//check connection database
	_, err := hdl.svc.GetAllUser(nil)
	if err != nil {
		response.Database = "Not OK"
	}

	return c.Status(200).JSON(response)
}
