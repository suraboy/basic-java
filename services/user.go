package services

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/suraboy/go-fiber-api/config"
	"github.com/suraboy/go-fiber-api/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

var messageError struct {
	Errors messageFormat `json:"errors"`
}

type messageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

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

func CreateNewProduct(c *fiber.Ctx) error {
	db := config.PostgresConnection()
	user := new(models.Users)
	var msgError messageFormat

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
