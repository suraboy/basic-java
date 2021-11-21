package http

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/internal/core/domain"
	responseBody "go-fiber-api/pkg/util"
	"net/http"
	"strconv"
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
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	var req domain.User
	if err := c.BodyParser(&req); err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	err = hdl.validator.ValidateStruct(req)
	if err != nil {
		return responseBody.SendValidationError(c,err)
	}

	response, err := hdl.svc.UpdateUserById(req,id)
	if err != nil {
		return responseBody.SendExceptionError(c,err)
	}
	return c.JSON(responseBody.RespondWithCollection(response))
}

func (hdl *Handler) DeleteUser(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	_, err = hdl.svc.DestroyUserById(id)
	if err != nil {
		return responseBody.SendExceptionError(c,err)
	}

	return c.Status(http.StatusNoContent).JSON("")
}
