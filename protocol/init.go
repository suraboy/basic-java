package protocol

import (
	"go-fiber-api/config"
	"go-fiber-api/infrastructure"
	"go-fiber-api/internal/core/service"
	"go-fiber-api/internal/repository/postgres"
	"go-fiber-api/pkg/logger"
	"go-fiber-api/pkg/validators"
)

var app *application

type application struct {
	// svc stand for service
	svc *service.Service
	// pkg stand for package
	pkg packages
}

type packages struct {
	validator validators.Validator
}

func init() {
	logger.Init()
	config.Init()
	packages := packages{
		validator: validators.NewValidator(),
	}

	// establish connections ...
	var (
		dbConn = infrastructure.PostgresConnection(infrastructure.Config{
			Username: config.GetViper().Database.Username,
			Password: config.GetViper().Database.Password,
			Host:     config.GetViper().Database.Host,
			Port:     config.GetViper().Database.Port,
			Database: config.GetViper().Database.Database,
		})
	)

	// repositories ...
	var (
		repo = postgres.NewPostgres(dbConn)
	)

	svc := service.New(service.Config{
		Repository: repo,
	})

	app = &application{
		svc: svc,
		pkg: packages,
	}
}
