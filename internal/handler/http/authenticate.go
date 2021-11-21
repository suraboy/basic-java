package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt/v4"
	"go-fiber-api/internal/core/domain"
	responseBody "go-fiber-api/pkg/util"
	"log"
	"time"
)

/**
 * Created by boy.sirichai on 21/11/2021 AD
 */

func (hdl *Handler) Authenticate(c *fiber.Ctx) error {
	var req domain.SignIn
	if err := c.BodyParser(&req); err != nil {
		return responseBody.SendBadRequest(c, err)
	}

	err := hdl.validator.ValidateStruct(req)
	if err != nil {
		return responseBody.SendValidationError(c, err)
	}
	// Throws Unauthorized error
	if req.Username != "admin" || req.Password != "admin" {
		return responseBody.SendUnauthorized(c)
	}

	// Create token
	res, err := generateAccessToken(req)

	if err != nil {
		return responseBody.SendInternalServerError(c, err)
	}

	return c.Status(200).JSON(res)
}

func generateAccessToken(req domain.SignIn) (domain.Token, error) {
	var res domain.Token
	// Create the Claims
	claims := jwt.MapClaims{
		"sub":   utils.UUID(),
		"name":  req.Username + " " + req.Password,
		"admin": true,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return res, err
	}

	res.AccessToken = accessToken

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = utils.UUID()
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt, err := refreshToken.SignedString([]byte("secret"))

	if err != nil {
		log.Printf("token.SignedString: %v", err)
		return res, err
	}
	res.RefreshToken = rt
	return res, nil
}
