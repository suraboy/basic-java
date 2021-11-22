package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-fiber-api/internal/core/domain"
	"go-fiber-api/pkg/logger"
	responseBody "go-fiber-api/pkg/util"
	"net/http"
	"strconv"
)

/**
 * Created by boy.sirichai on 22/11/2021 AD
 */

func (hdl *Handler) GetAllProduct(c *fiber.Ctx) error {
	var req domain.Product
	if err := c.QueryParser(&req); err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	response, err := hdl.svc.GetAllProduct(req)
	if err != nil {
		return responseBody.SendNotFound(c)
	}

	return c.JSON(responseBody.RespondWithCollection(response))
}


func (hdl *Handler) FindProductById(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	response, err := hdl.svc.FindProductById(domain.Product{ID: uint(id)})
	if err != nil {
		return responseBody.SendNotFound(c)
	}

	return c.JSON(responseBody.RespondWithCollection(response))
}

func (hdl *Handler) CreateProduct(c *fiber.Ctx) error {
	var req domain.Product

	if err := c.BodyParser(&req); err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	err := hdl.validator.ValidateStruct(req)
	if err != nil {
		return responseBody.SendValidationError(c,err)
	}

	user := c.Locals("user").(*jwt.Token)
	logger.Infof(`Token: %s`,user)
	claims := user.Claims.(jwt.MapClaims)
	uuid := claims["sub"].(string)
	req.CreatedBy = uuid
	req.UpdatedBy = uuid

	response, err := hdl.svc.CreateProduct(req)
	if err != nil {
		return err
	}

	return c.JSON(responseBody.RespondWithCollection(response))
}

func (hdl *Handler) UpdateProduct(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	var req domain.Product
	if err := c.BodyParser(&req); err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	err = hdl.validator.ValidateStruct(req)
	if err != nil {
		return responseBody.SendValidationError(c,err)
	}

	response, err := hdl.svc.UpdateProductById(req,id)
	if err != nil {
		return responseBody.SendExceptionError(c,err)
	}
	return c.JSON(responseBody.RespondWithCollection(response))
}

func (hdl *Handler) DeleteProduct(c *fiber.Ctx) error {
	paramID := c.Params("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return responseBody.SendBadRequest(c,err)
	}

	_, err = hdl.svc.DestroyProductById(domain.Product{ID: uint(id)})
	if err != nil {
		return responseBody.SendExceptionError(c,err)
	}

	return c.Status(http.StatusNoContent).JSON("")
}