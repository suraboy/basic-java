package infrastructure

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/suraboy/go-fiber-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
)


func PostgresConnection() *gorm.DB {
	log.Print("Connecting to PostgreSQL DB...")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v TimeZone=Asia/Bangkok",
		config.GetViper().Database.Host,
		config.GetViper().Database.Username,
		config.GetViper().Database.Password,
		config.GetViper().Database.Database,
		config.GetViper().Database.Port,
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
