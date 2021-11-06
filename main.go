package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/suraboy/go-fiber-api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("$PORT must be set : %s", port)
	}

	router := fiber.New()
	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	//check version project
	router.Get("/version", func(c *fiber.Ctx) error {
		return c.Send([]byte(os.Getenv("TAG_VERSION")))
	})

	router.Static("/", "./storage/images", fiber.Static{
		Index: "",
	})

	routes.UserV1Route(router)
	routes.AuthenticateV1Route(router)
	routes.ProductV1Route(router)

	if err := router.Listen(":" + port); err != nil {
		log.Fatalf("shutting down the server : %s", err)
	}
}
