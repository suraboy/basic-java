package port

import "github.com/suraboy/go-fiber-api/internal/core/domain/models"

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
	GetAllUser(request *models.User) ([]models.User, error)
	FindUserById(request *models.User) (models.User, error)
}
