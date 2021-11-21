package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go-fiber-api/internal/core/domain"
	"go-fiber-api/pkg/logger"
	responseBody "go-fiber-api/pkg/util"
	"golang.org/x/crypto/bcrypt"
	"time"
)

/**
 * Created by boy.sirichai on 21/11/2021 AD
 */
var c *fiber.Ctx

func (s Service) SignIn(req domain.SignIn) (domain.Token, error) {
	var user domain.User
	user.Username = req.Username
	userInfo, err := s.repository.FindUser(user)
	if err != nil {
		return domain.Token{}, responseBody.SendNotFound()
	}
	//compare password
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password))
	if err != nil {
		return domain.Token{}, responseBody.SendUnauthorized()
	}
	//generate token
	res, err := GenerateAccessToken(userInfo)

	if err != nil {
		return domain.Token{}, responseBody.SendInternalServerError(err)
	}

	return res, nil
}

func (s Service) RefreshToken(req domain.RefreshToken) (domain.Token, error) {
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("secret"), nil
	})

	if err != nil {
		return domain.Token{}, responseBody.SendBadRequest(err)
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		uid, ok := claims["sub"].(string)
		if !ok {
			return domain.Token{}, responseBody.SendBadRequest(err)
		}
		var user domain.User
		user.Uuid = uid

		userInfo, err := s.repository.FindUser(user)

		if err != nil {
			return domain.Token{}, responseBody.SendNotFound()
		}

		res, err := GenerateAccessToken(userInfo)

		if err != nil {
			return domain.Token{}, responseBody.SendInternalServerError(err)
		}
		return res, nil
	} else {
		return domain.Token{}, responseBody.SendTokenExpired()
	}
}

func GenerateAccessToken(req domain.User) (domain.Token, error) {
	var res domain.Token
	// Create the Claims
	claims := jwt.MapClaims{
		"sub":  req.Uuid,
		"name": req.Name + " " + req.LastName,
		"type": req.Type,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	accessToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		logger.Infof("token.SignedString: %v", err)
		return res, err
	}

	res.AccessToken = accessToken

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = req.Uuid
	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	rt, err := refreshToken.SignedString([]byte("secret"))

	if err != nil {
		logger.Infof("token.SignedString: %v", err)
		return res, err
	}
	res.RefreshToken = rt
	return res, nil
}
