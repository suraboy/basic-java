package http

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/internal/core/domain"
	"go-fiber-api/pkg/logger"
	responseBody "go-fiber-api/pkg/util"
	"net/http"
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
		return responseBody.SendBadRequest(c,err)
	}

	response, err := hdl.svc.GetAllUser(req)
	if err != nil {
		return responseBody.SendNotFound(c)
	}

	return c.JSON(responseBody.RespondWithCollection(response))
}

func (hdl *Handler) FindUserById(c *fiber.Ctx) error {
	var req domain.User
	if err := c.QueryParser(&req); err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	response, err := hdl.svc.FindUserById(req)
	if err != nil {
		return responseBody.SendNotFound(c)
	}

	return c.JSON(responseBody.RespondWithCollection(response))
}

func (hdl *Handler) CreateUser(c *fiber.Ctx) error {
	var req domain.User
	if err := c.BodyParser(&req); err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	err := hdl.validator.ValidateStruct(req)
	if err != nil {
		return responseBody.SendValidationError(c,err)
	}

	response, err := hdl.svc.CreateUser(req)
	if err != nil {
		return err
	}

	return c.JSON(responseBody.RespondWithCollection(response))
}

func (hdl *Handler) UpdateUser(c *fiber.Ctx) error {
	var req domain.User
	logger.Infof("Request0 : %s",req )
	if err := c.BodyParser(&req); err != nil {
		logger.Infof("Request1 : %s",req )
		return responseBody.SendBadRequest(c,err)
	}

	logger.Infof("Request2 : %s",req )
	err := hdl.validator.ValidateStruct(req)
	if err != nil {
		return responseBody.SendValidationError(c,err)
	}

	response, err := hdl.svc.UpdateUserById(req)
	if err != nil {
		return err
	}
	return c.JSON(responseBody.RespondWithCollection(response))
}

func (hdl *Handler) DeleteUser(c *fiber.Ctx) error {
	var req domain.User
	id := c.Params("id")
	if err := c.QueryParser(&req); err != nil {
		return c.JSON(responseBody.SendBadRequest(c,err))
	}

	_, err := hdl.svc.DestroyUserById(req, id)
	if err != nil {
		return err
	}

	return c.Status(http.StatusNoContent).JSON("")
}
