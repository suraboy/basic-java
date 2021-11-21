package service

import (
	"go-fiber-api/internal/core/domain/models"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

func (s Service) GetAllUser(request *models.User) ([]models.User, error) {
	return s.repository.GetUsers(request)
}

func (s Service) FindUserById(request *models.User) (models.User, error) {
	return s.repository.FindUser(request)
}
