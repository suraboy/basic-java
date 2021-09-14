package config

import (
	"database/sql"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func PostgresConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}

	log.Print("Connecting to PostgreSQL DB...")

	sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	DB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("connected")

	return DB
}
