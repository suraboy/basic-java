package service

import (
	"go-fiber-api/internal/core/domain"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

func (s Service) GetAllUser(request *domain.User) ([]domain.User, error) {
	return s.repository.GetUsers(request)
}

func (s Service) FindUserById(request *domain.User) (domain.User, error) {
	return s.repository.FindUser(request)
}
