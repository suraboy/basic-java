package service

import (
	"github.com/gofiber/fiber/v2/utils"
	"go-fiber-api/internal/core/domain"
	responseBody "go-fiber-api/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

func (s Service) GetAllUser(request domain.User) ([]domain.User, error) {
	return s.repository.GetUsers(request)
}

func (s Service) FindUserById(request domain.User) (domain.User, error) {
	return s.repository.FindUser(request)
}

func (s Service) CreateUser(request domain.User) (domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, responseBody.SendInternalServerError(c, err)
	}
	request.Uuid = utils.UUID()
	request.Password = string(hashedPassword)
	res, err := s.repository.CreateUser(request)

	if err != nil {
		return domain.User{}, responseBody.SendExceptionError(c, err)
	}
	return res, nil
}

func (s Service) UpdateUserById(request domain.User, id int) (domain.User, error) {
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return domain.User{}, err
		}
		request.Password = string(hashedPassword)
	}

	_, err := s.repository.UpdateUser(request, id)
	if err != nil {
		return domain.User{}, err
	}
	userInfo, err := s.repository.FindUser(domain.User{ID: uint(id)})
	if err != nil {
		return domain.User{}, err
	}

	return userInfo, nil
}

func (s Service) DestroyUserById(id int) (domain.User, error) {
	_, err := s.repository.FindUser(domain.User{ID: uint(id)})
	if err != nil {
		return domain.User{}, err
	}

	res, err := s.repository.DeleteUser(id)
	if err != nil {
		return domain.User{}, responseBody.SendExceptionError(c, err)
	}
	return res, nil
}
