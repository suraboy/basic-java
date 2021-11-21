package infrastructure

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
)


type Config struct {
	Username string
	Password string
	Host     string
	Port     int64
	Database string
}

func PostgresConnection(config Config) *gorm.DB {
	log.Print("Connecting to PostgreSQL DB...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v TimeZone=Asia/Bangkok",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		logrus.Fatalf("Failed to connect to database : %s", err)
	}


	_ = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{postgres.Open(dsn)},
		Replicas: []gorm.Dialector{postgres.Open(dsn)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	logrus.Infof("connect postgres success : %s", dsn)
	log.Println("connected")
	return db
}
