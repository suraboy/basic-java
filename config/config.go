package config

import (
	"log"

	"github.com/suraboy/go-fiber-api/pkg/logger"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ",squash" will ignore mapstructure for that field

type Config struct {
	App      appConfig      `mapstructure:",squash"`
	Database databaseConfig `mapstructure:",squash"`
}

type appConfig struct {
	HTTPPort string `mapstructure:"APP_HTTP_PORT"`
	Env      string `mapstructure:"APP_ENV"`
}

type databaseConfig struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     int64  `mapstructure:"POSTGRES_PORT"`
	Database string `mapstructure:"POSTGRES_DBNAME"`
	Username string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
}

var config Config

func Init(cfgPath string) {

	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.AppLogger.Infof("Config file has changed: %s", e.Name)
	})
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetViper() *Config {
	return &config
}
