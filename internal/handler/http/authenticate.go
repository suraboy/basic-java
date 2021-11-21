package http

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-api/internal/core/domain"
	responseBody "go-fiber-api/pkg/util"
)

/**
 * Created by boy.sirichai on 21/11/2021 AD
 */

func (hdl *Handler) Authenticate(c *fiber.Ctx) error {
	var req domain.SignIn
	if err := c.BodyParser(&req); err != nil {
		return responseBody.SendBadRequest(err)
	}

	err := hdl.validator.ValidateStruct(req)
	if err != nil {
		return responseBody.SendValidationError(err)
	}
	//check user exits
	res, err := hdl.svc.SignIn(req)

	if err != nil {
		return err
	}
	return c.Status(200).JSON(res)
}

func (hdl *Handler) RefreshToken(c *fiber.Ctx) error {
	var req domain.RefreshToken
	if err := c.BodyParser(&req); err != nil {
		return responseBody.SendBadRequest(err)
	}

	err := hdl.validator.ValidateStruct(req)
	if err != nil {
		return responseBody.SendValidationError(err)
	}

	res, err := hdl.svc.RefreshToken(req)

	if err != nil {
		return err
	}

	return c.Status(200).JSON(res)
}

func (hdl *Handler) Logout(c *fiber.Ctx) error {
	return c.JSON(responseBody.CustomResponse(200,"You must handler redis for remove token"))
}
