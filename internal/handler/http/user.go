package http

import (
	"github.com/gofiber/fiber/v2"
	domain "github.com/suraboy/go-fiber-api/internal/core/domain/models"
	transform "github.com/suraboy/go-fiber-api/pkg/util"
	"github.com/suraboy/go-fiber-api/pkg/validators"
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

// @Tags api/kkbank
// @Success 200
// @Method GET
// @Router /api/kkbank/inquiry [GET]

func (hdl *HTTPHandler) GetAllUser(c *fiber.Ctx) error {
	var req domain.User
	if err := c.QueryParser(&req); err != nil {
		return c.JSON(transform.SendBadRequest(c,err))
	}

	validate := validators.NewValidator()
	err := validate.ValidateStruct(req)
	if err != nil {
		return c.JSON(transform.SendValidationError(c,err))
	}

	response, err := hdl.svc.GetAllUser(&req)
	if err != nil {
		return c.JSON(transform.SendNotFound(c))
	}

	return c.JSON(transform.RespondWithCollection(response))
}
