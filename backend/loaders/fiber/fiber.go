package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Init() {
	// Initialize fiber instance
	app := fiber.New(fiber.Config{
		ErrorHandler:  errorHandler,
		Prefork:       false,
		StrictRouting: true,
		ServerHeader:  config.C.ServerHeader,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
	})

	// Init root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(responder.InfoResponse{
			Success: true,
			Message: "EGG_FACTORY_API_ROOT",
		})
	})

	// * Init swagger endpoint
	swaggerGroup := app.Group("swagger/")
	if config.C.Environment.String() == "dev" {
		swagger.Init(swaggerGroup)
	}

	// * Init websocket
	websocketGroup := app.Group("ws/")
	websocket.Init(websocketGroup)

	// Init API endpoints
	apiGroup := app.Group("api/")

	apiGroup.Use(middlewares.Limiter)
	apiGroup.Use(middlewares.Cors)
	apiGroup.Use(middlewares.Recover)

	endpoints.Init(apiGroup)

	// Init not found handler
	app.Use(notfoundHandler)

	// Startup
	err := app.Listen(config.C.Address)
	if err != nil {
		logger.Log(logrus.Fatal, err.Error())
	}
}
