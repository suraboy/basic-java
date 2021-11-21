package http

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/internal/core/domain"
	transform "go-fiber-api/pkg/util"
)

/*
	|--------------------------------------------------------------------------
	| The Handler Adaptor
	|--------------------------------------------------------------------------
	|
	| An Adapter will initiate the interaction with the Application through
	| a Port, using specific technology that means you can choose
	| any technology you want for your application protocol.
	|
*/

func (hdl *Handler) GetAllUser(c *fiber.Ctx) error {
	var req domain.User
	if err := c.QueryParser(&req); err != nil {
		return transform.SendBadRequest(c,err)
	}

	response, err := hdl.svc.GetAllUser(&req)
	if err != nil {
		return transform.SendNotFound(c)
	}

	return c.JSON(transform.RespondWithCollection(response))
}

func (hdl *Handler) FindUserById(c *fiber.Ctx) error {
	var req domain.User
	if err := c.QueryParser(&req); err != nil {
		return c.JSON(transform.SendBadRequest(c,err))
	}

	response, err := hdl.svc.FindUserById(&req)
	if err != nil {
		return transform.SendNotFound(c)
	}

	return c.JSON(transform.RespondWithCollection(response))
}
