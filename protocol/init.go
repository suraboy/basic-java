package protocol

import (
	"github.com/suraboy/go-fiber-api/config"
	"github.com/suraboy/go-fiber-api/internal/core/service"
	"github.com/suraboy/go-fiber-api/internal/repository/postgres"
	"github.com/suraboy/go-fiber-api/pkg/logger"
	"github.com/suraboy/go-fiber-api/pkg/validators"
)

// export `CfgPath` to set config path from cmd/root.go
var CfgPath string

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
	config.Init(CfgPath)
	packages := packages{
		validator: validators.NewValidator(),
	}

	repo := postgres.NewPostgres()

	app = &application{
		svc: service.New(repo),
		pkg: packages,
	}
}
