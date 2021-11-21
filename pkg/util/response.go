package util

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var MessageError struct {
	Errors MessageFormat `json:"errors"`
}

type MessageFormat struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
	Error      string `json:"error,omitempty"`
}

type ResponseCollection struct {
	Data interface{} `json:"data"`
}

var c *fiber.Ctx

func RespondWithCollection(result interface{}) *ResponseCollection {
	if result == nil {
		type Empty struct{}
		result = Empty{}
	}
	return &ResponseCollection{
		Data: result,
	}
}

func CustomResponse(httpCode int, message string) MessageFormat {
	var msgError MessageFormat
	msgError.StatusCode = httpCode
	msgError.Message = message
	MessageError.Errors = msgError
	return msgError
}

func SendNotFound() error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusNotFound
	msgError.Message = "Not Found"
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendUnauthorized() error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusUnauthorized
	msgError.Message = "Unauthorized"
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendTokenExpired() error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusUnauthorized
	msgError.Message = "Token Expired"
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendInternalServerError(err error) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusInternalServerError
	msgError.Message = "Internal Server Error"
	msgError.Error = err.Error()
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendBadRequest(err error) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusBadRequest
	msgError.Message = "Bad Request"
	msgError.Error = err.Error()
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}

func SendValidationError(err error) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusUnprocessableEntity
	msgError.Message = "The given data was invalid."
	msgError.Error = err.Error()
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}


func SendExceptionError(err error) error {
	var msgError MessageFormat
	msgError.StatusCode = http.StatusExpectationFailed
	msgError.Message = "Expectation Failed"
	msgError.Error = err.Error()
	MessageError.Errors = msgError

	return c.Status(MessageError.Errors.StatusCode).JSON(MessageError)
}
