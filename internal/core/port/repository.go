package port

import "go-fiber-api/internal/core/domain/models"

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
	GetUsers(request *models.User) ([]models.User, error)
	FindUser(request *models.User) (models.User, error)
}
