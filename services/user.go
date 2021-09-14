package services

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/suraboy/go-fiber-api/config"
	"github.com/suraboy/go-fiber-api/models"
	"github.com/suraboy/go-fiber-api/utils"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func GetAllUser(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	var user []models.Users
	db.Find(&user)
	return c.JSON(fiber.Map{"data": user})
}

func FindUser(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	id := c.Params("id")
	user := models.Users{}

	if err := db.Find(&user, id).Error; err != nil || db.Find(&user, id).RowsAffected == 0 {
		var msgError messageFormat
		if db.Find(&user, id).RowsAffected == 0 {
			msgError.StatusCode = http.StatusNotFound
			msgError.Message = "Not Found"
		} else {
			msgError.StatusCode = http.StatusInternalServerError
			msgError.Message = "Internal Server Error"
			msgError.Error = err.Error()
		}
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	user := new(models.Users)
	var msgError messageFormat

	if err := c.BodyParser(user); err != nil {
		msgError.StatusCode = http.StatusBadRequest
		msgError.Message = "Bad Request"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		msgError.StatusCode = http.StatusUnprocessableEntity
		msgError.Message = "The given data was invalid."
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.Password = string(hashedPassword)
	if err := db.Create(&user).Error; err != nil {
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{"data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	pass := ""
	db := config.PostgresConnection()
	id := c.Params("id")
	user := models.Users{}
	var msgError messageFormat

	if err := c.BodyParser(&user); err != nil {
		msgError.StatusCode = http.StatusBadRequest
		msgError.Message = "Bad Request"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}

	if user.Password != "" {
		password := []byte(user.Password)
		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		pass = string(hashedPassword)
	}

	if db.Find(&user, id).RowsAffected == 0 {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}

	if pass != "" {
		user.Password = pass
	}

	if err := db.Save(&user).Error; err != nil {
		msgError.StatusCode = http.StatusExpectationFailed
		msgError.Message = "Expectation Failed"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"data": user})
}

func DeleteUser(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	db := config.PostgresConnection()
	user := models.Users{}
	var msgError messageFormat
	if db.Find(&user, id).RowsAffected == 0 {
		msgError.StatusCode = http.StatusNotFound
		msgError.Message = "Not Found"
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}

	if err := db.Delete(&user).Error; err != nil {
		msgError.StatusCode = http.StatusInternalServerError
		msgError.Message = "Internal Server Error"
		msgError.Error = err.Error()
		messageError.Errors = msgError
		return c.Status(msgError.StatusCode).JSON(messageError)
	}

	return c.Status(http.StatusNoContent).JSON(nil)
}
