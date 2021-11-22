package protocol

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"
	"go-fiber-api/config"
	"go-fiber-api/internal/handler/http"
	route "go-fiber-api/internal/routes"
	"log"
)

/*
	|--------------------------------------------------------------------------
	| Application Protocol
	|--------------------------------------------------------------------------
	|
	| Here you can choose which protocol your application wants to interact
	| with the client for instance HTTP, gRPC etc.
	|
*/

func ServeREST() error {
	f := fiber.New(fiber.Config{
		DisableKeepalive: false,
	})

	f.Use(recovery.New())
	f.Use(cors.New())

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	hdl := http.NewHandler(app.svc, app.pkg.validator)
	//health-check project
	f.Get("/healthcheck", hdl.HealthCheck)
	//check version project
	f.Get("/version", hdl.CheckVersion)

	v1 := f.Group("/v1")
	v1.Static("/", "./storage/images", fiber.Static{
		Index: "",
	})

	route.AuthV1Route(v1, hdl)
	route.UserV1Route(v1, hdl)
	route.ProductV1Route(v1, hdl)

	if err := f.Listen(":" + config.GetViper().App.Port); err != nil {
		log.Fatalf("shutting down the server : %s", err)
	}

	return nil
}
