package protocol

import (
	"github.com/suraboy/go-fiber-api/internal/handler/http"
	"github.com/suraboy/go-fiber-api/internal/route"
	"github.com/suraboy/go-fiber-api/pkg/logger"
	"github.com/suraboy/go-fiber-api/pkg/middleware"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	recovery "github.com/gofiber/fiber/v2/middleware/recover"
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

	middleware := middleware.NewMiddleware()

	f.Use(middleware.LogRequest)
	f.Use(recovery.New())
	f.Use(cors.New())

	hdl := http.NewHTTPHandler(app.svc, app.pkg.validator)

	v1 := f.Group("/v1")
	v1.Get("/healthcheck", hdl.HealthCheck)
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	//check version project
	v1.Get("/version", func(c *fiber.Ctx) error {
		return c.Send([]byte(os.Getenv("TAG_VERSION")))
	})

	v1.Static("/", "./storage/images", fiber.Static{
		Index: "",
	})

	route.UserV1Route(v1, hdl)

	err := f.Listen(":8080")
	if err != nil {
		return err
	}
	// gracefully shuts down  ...
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		// Block until got a signal.
		<-c
		logger.AppLogger.Info("Gracefully shutting down...")
		err = f.Shutdown()
		if err != nil {
			logger.AppLogger.Infof("Got error: %s while shutting down HTTP server", err)
		}
		os.Exit(0)
	}()

	return nil
}
