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

type Repository interface {
	GetAllUser(request *models.User, field ...interface{}) (models.User, error)
}
