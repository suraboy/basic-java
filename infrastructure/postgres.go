package infrastructure

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
)


func PostgresConnection() *gorm.DB {
	log.Print("Connecting to PostgreSQL DB...")

	dsn := "host=ec2-107-20-24-247.compute-1.amazonaws.com user=tfmdhiaxecdrzy password=4d65da8a1ed2426b96429138ed9a62029712f0b6d6f737505f4ff1e422508a54 dbname=d6gg7trl9p9gub port=5432 TimeZone=Asia/Bangkok"
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

	logrus.Infof("connect mysql success : %s", dsn)
	log.Println("connected")
	return db
}
