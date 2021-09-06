package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
)

func connection(dsn string, source string, replica string) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	conn.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{postgres.Open(source)},
		Replicas: []gorm.Dialector{postgres.Open(replica)},
		Policy:   dbresolver.RandomPolicy{},
	}))
	if err != nil {
		log.Fatalf("cannot open postgres connection:%s", err)
	}

	log.Fatalf("postgres connection:%s", conn)

	return conn
}

//PostgresConnection connect postgres db
func PostgresConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file:%s", err)
	}
	DSN := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DBNAME"))

	source := fmt.Sprintf("user=%s password=%s host=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"))

	replica := fmt.Sprintf("user=%s password=%s host=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST_REPLI1"))

	return connection(DSN, source, replica)
}
