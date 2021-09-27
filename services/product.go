package services

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/suraboy/go-fiber-api/config"
	"github.com/suraboy/go-fiber-api/models"
	"github.com/suraboy/go-fiber-api/utils"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

func GetAllProduct(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	var user []models.Products
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))
	db.Limit(limit).Offset(offset).Find(&user)
	return c.JSON(fiber.Map{"data": user})
}

func FindProduct(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	id := c.Params("id")
	user := models.Products{}
	response := db.Find(&user, id)
	if err := response.Error; err != nil || response.RowsAffected == 0 {
		if response.RowsAffected == 0 {
			return utils.SendNotFound(c)
		} else {
			return utils.SendInternalServerError(c, err)
		}
	}
	return c.JSON(user)
}

func CreateProduct(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	user := new(models.Products)
	if err := c.BodyParser(user); err != nil {
		return utils.SendBadRequest(c, err)
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return utils.SendValidationError(c, err)
	}
	if err := db.Create(&user).Error; err != nil {
		return utils.SendExceptionError(c, err)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": user})
}

func UpdateProduct(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	id := c.Params("id")
	user := models.Products{}
	if db.Find(&user, id).RowsAffected == 0 {
		return utils.SendNotFound(c)
	}

	if err := c.BodyParser(&user); err != nil {
		return utils.SendBadRequest(c, err)
	}

	if err := db.Save(&user).Error; err != nil {
		return utils.SendExceptionError(c, err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": user})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.PostgresConnection()
	user := models.Products{}
	if db.Find(&user, id).RowsAffected == 0 {
		return utils.SendNotFound(c)
	}

	if err := db.Delete(&user).Error; err != nil {
		return utils.SendInternalServerError(c, err)
	}

	return c.Status(http.StatusNoContent).JSON(nil)
}
