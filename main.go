package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/suraboy/go-fiber-api/routes"
)

const AppVersion = "1.0.1"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := fiber.New()
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	//check version project
	router.Get("/version", func(c *fiber.Ctx) error {
		return c.Send([]byte(AppVersion))
	})

	routes.UserV1Route(router)

	if err := router.Listen(":" + port); err != nil {
		log.Fatalf("shutting down the server : %s", err)
	}
}
