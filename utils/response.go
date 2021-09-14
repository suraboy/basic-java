package utils

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var MessageError struct {
	Errors MessageFormat `json:"errors"`
}

type MessageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

func CustomResponse(httpCode int, message string, error string) MessageFormat {
	var msgError MessageFormat
	msgError.StatusCode = httpCode
	msgError.Message = message
	msgError.Error = error
	MessageError.Errors = msgError
	return msgError
}

func SendNotFound(c *fiber.Ctx) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusNotFound
	msgError.Message = "Not Found"
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendInternalServerError(c *fiber.Ctx,err error) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusInternalServerError
	msgError.Message = "Internal Server Error"
	msgError.Error = err.Error()
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendBadRequest(c *fiber.Ctx,err error) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusBadRequest
	msgError.Message = "Bad Request"
	msgError.Error = err.Error()
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendValidationError(c *fiber.Ctx,err error) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusUnprocessableEntity
	msgError.Message = "The given data was invalid."
	msgError.Error = err.Error()
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendExceptionError(c *fiber.Ctx,err error) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusExpectationFailed
	msgError.Message = "Expectation Failed"
	msgError.Error = err.Error()
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}
