package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"strings"

	"go-fiber-api/pkg/logger"
)

// ",squash" will ignore envconfig for that field

type Config struct {
	App      appConfig
	Database databaseConfig
}

type appConfig struct {
	HTTPPort string `envconfig:"APP_HTTP_PORT"`
	Env      string `envconfig:"APP_ENV"`
	APIKey   string `envconfig:"APP_API_KEY"`
}

type databaseConfig struct {
	Host     string `envconfig:"POSTGRES_HOST"`
	Port     int64  `envconfig:"POSTGRES_PORT"`
	Database string `envconfig:"POSTGRES_DBNAME"`
	Username string `envconfig:"POSTGRES_USER"`
	Password string `envconfig:"POSTGRES_PASSWORD"`
}

var config Config

// Init is application config initialization ...
func Init() {
	err := godotenv.Load()
	if err != nil {
		envFileNotFound := strings.Contains(err.Error(), "no such file or directory")
		if !envFileNotFound {
			logger.Infof("read config error: %v", err)
		} else {
			logger.Info("use environment from OS")
		}
	}
	err = envconfig.Process("", &config)
	if err != nil {
		logger.Infof("parse config error: %v", err)
	}

}

func GetViper() *Config {
	return &config
}
