package service

import (
	"github.com/suraboy/go-fiber-api/internal/core/port"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

type Service struct {
	repository       port.Repository
}

func New(repo port.Repository) *Service {
	return &Service{
		repository:   repo,
	}
}
