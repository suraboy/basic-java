package services

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/suraboy/go-fiber-api/config"
	"github.com/suraboy/go-fiber-api/pkg"
	"github.com/suraboy/go-fiber-api/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

func GetAllUser(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	var user []models.User
	limit, _ := strconv.Atoi(c.Query("limit"))
	offset, _ := strconv.Atoi(c.Query("offset"))

	db.Limit(limit).Offset(offset).Find(&user)
	return c.JSON(fiber.Map{"data": user})
}

func FindUser(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	id := c.Params("id")
	user := models.User{}
	response := db.Find(&user, id)
	if err := response.Error; err != nil || response.RowsAffected == 0 {
		if response.RowsAffected == 0 {
			return pkg.SendNotFound(c)
		} else {
			return pkg.SendInternalServerError(c, err)
		}
	}
	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return pkg.SendBadRequest(c, err)
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return pkg.SendValidationError(c, err)
	}
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		return pkg.SendExceptionError(c, err)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	pass := ""
	db := config.PostgresConnection()
	id := c.Params("id")
	user := models.User{}
	if db.Find(&user, id).RowsAffected == 0 {
		return pkg.SendNotFound(c)
	}

	if err := c.BodyParser(&user); err != nil {
		return pkg.SendBadRequest(c, err)
	}

	if user.Password != "" {
		password := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		pass = string(hashedPassword)
	}

	if pass != "" {
		user.Password = pass
	}

	if err := db.Save(&user).Error; err != nil {
		return pkg.SendExceptionError(c, err)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": user})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.PostgresConnection()
	user := models.User{}
	if db.Find(&user, id).RowsAffected == 0 {
		return pkg.SendNotFound(c)
	}

	if err := db.Delete(&user).Error; err != nil {
		return pkg.SendInternalServerError(c, err)
	}

	return c.Status(http.StatusNoContent).JSON(nil)
}
