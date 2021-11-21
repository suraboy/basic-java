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

type Repository interface {
	// GetUsers FindUser CreateUser UpdateUser DeleteUser feature: users
	GetUsers(request domain.User) ([]domain.User, error)
	FindUser(request domain.User) (domain.User, error)
	CreateUser(request domain.User) (domain.User, error)
	UpdateUser(request domain.User) (domain.User, error)
	DeleteUser(request domain.User, id string) (domain.User, error)
}
