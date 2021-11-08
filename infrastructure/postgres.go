package infrastructure

import (
	"database/sql"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)


func PostgresConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}

	log.Print("Connecting to PostgreSQL DB...")

	sqlDB, err := sql.Open("postgres", "postgres://tfmdhiaxecdrzy:4d65da8a1ed2426b96429138ed9a62029712f0b6d6f737505f4ff1e422508a54@ec2-107-20-24-247.compute-1.amazonaws.com:5432/d6gg7trl9p9gub")
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	return gormDB
}
