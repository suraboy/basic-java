package service

import (
	"go-fiber-api/internal/core/port"
)

/*
	|--------------------------------------------------------------------------
	| Application's Business Logic
	|--------------------------------------------------------------------------
	|
	| Here you can implement a business logic  for your application
	|
*/

type Config struct {
	Repository port.Repository
}

type Service struct {
	repository port.Repository
}

func New(cfg Config) *Service {
	return &Service{
		repository: cfg.Repository,
	}
}
