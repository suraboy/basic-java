package protocol

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
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

func ServeREST() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("$PORT must be set : %s", port)
	}

	f := fiber.New()
	//
	//middleware := middleware.NewMiddleware()
	//
	//f.Use(middleware.LogRequest)
	//f.Use(recovery.New())
	//f.Use(cors.New())

	v1 := f.Group("/v1")
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
	//
	//hdl := http.NewHTTPHandler(app.svc, app.pkg.validator)
	//v1.Get("/healthcheck", hdl.HealthCheck)
	////route.UserV1Route(v1, hdl)

	if err := f.Listen(":" + port); err != nil {
		log.Fatalf("shutting down the server : %s", err)
	}
	//// gracefully shuts down  ...
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//
	//go func() {
	//	// Block until got a signal.
	//	<-c
	//	logger.AppLogger.Info("Gracefully shutting down...")
	//	err = f.Shutdown()
	//	if err != nil {
	//		logger.AppLogger.Infof("Got error: %s while shutting down HTTP server", err)
	//	}
	//	os.Exit(0)
	//}()
	//
	//return nil
}
