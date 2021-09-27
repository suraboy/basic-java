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
	var product []models.Product
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	db.Limit(limit).Offset(offset).Find(&product)
	return c.JSON(fiber.Map{"data": product})
}

func FindProduct(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	id := c.Params("id")
	product := models.Product{}
	response := db.Find(&product, id)

	if err := response.Error; err != nil || response.RowsAffected == 0 {
		if response.RowsAffected == 0 {
			return utils.SendNotFound(c)
		} else {
			return utils.SendInternalServerError(c, err)
		}
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return utils.SendBadRequest(c, err)
	}

	validate := validator.New()
	err := validate.Struct(product)

	if err != nil {
		return utils.SendValidationError(c, err)
	}

	if err := db.Create(&product).Error; err != nil {
		return utils.SendExceptionError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": product})
}

func UpdateProduct(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	id := c.Params("id")
	product := models.Product{}

	if db.Find(&product, id).RowsAffected == 0 {
		return utils.SendNotFound(c)
	}

	if err := c.BodyParser(&product); err != nil {
		return utils.SendBadRequest(c, err)
	}

	if err := db.Save(&product).Error; err != nil {
		return utils.SendExceptionError(c, err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": product})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.PostgresConnection()

	product := models.Product{}

	if db.Find(&product, id).RowsAffected == 0 {
		return utils.SendNotFound(c)
	}

	if err := db.Delete(&product).Error; err != nil {
		return utils.SendInternalServerError(c, err)
	}

	return c.Status(http.StatusNoContent).JSON(nil)
}
