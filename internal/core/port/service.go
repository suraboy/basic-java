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
	//feature : authenticate
	SignIn(req domain.SignIn) (domain.Token, error)
	RefreshToken(req domain.RefreshToken) (domain.Token, error)
	//feature : users
	GetAllUser(request domain.User) ([]domain.User, error)
	FindUserById(request domain.User) (domain.User, error)
	CreateUser(request domain.User) (domain.User, error)
	UpdateUserById(request domain.User,id int) (domain.User, error)
	DestroyUserById(request domain.User) (domain.User, error)
	//feature : products
	GetAllProduct(request domain.Product) ([]domain.Product, error)
	FindProductById(request domain.Product) (domain.Product, error)
	CreateProduct(request domain.Product) (domain.Product, error)
	UpdateProductById(request domain.Product,id int) (domain.Product, error)
	DestroyProductById(request domain.Product) (domain.Product, error)
}
