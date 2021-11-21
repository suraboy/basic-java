package port

import (
	"go-fiber-api/internal/core/domain"
)

/*
	|--------------------------------------------------------------------------
	| Application Port
	|--------------------------------------------------------------------------
	|
	| Here you can define an interface which will allow foreign actors to
	| communicate with the Application
	|
*/

type Service interface {
	// SignIn RefreshToken feature : authenticate
	SignIn(req domain.SignIn) (domain.Token, error)
	RefreshToken(req domain.RefreshToken) (domain.Token, error)

	// GetAllUser FindUserById  feature : users
	GetAllUser(request domain.User) ([]domain.User, error)
	FindUserById(request domain.User) (domain.User, error)
	CreateUser(request domain.User) (domain.User, error)
	UpdateUserById(request domain.User) (domain.User, error)
	DestroyUserById(request domain.User, id string) (domain.User, error)
}
