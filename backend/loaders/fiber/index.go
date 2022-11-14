package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"backend/loaders/fiber/middleware"
	"backend/loaders/websocket"
	"backend/router"
	"backend/types/response"
	"backend/utils/config"
)

var app *fiber.App

func Init() {
	// Initialize fiber instance
	app = fiber.New(fiber.Config{
		ErrorHandler:  errorHandler,
		Prefork:       false,
		StrictRouting: true,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  5 * time.Second,
		AppName:       "ICT CHALLENGE API",
	})

	// Register root endpoint
	app.All("/", func(c *fiber.Ctx) error {
		return c.JSON(response.InfoResponse{
			Success: true,
			Message: "ICT CHALLENGE API ROOT",
		})
	})

	// Register API endpoints
	apiGroup := app.Group("api/")
	apiGroup.Use(middleware.Limiter)
	apiGroup.Use(middleware.Cors)
	router.Init(apiGroup)

	wsGroup := app.Group("ws/")
	websocket.Init(wsGroup)

	// Register not found handler
	app.Use(notfoundHandler)

	// Startup
	if err := app.Listen(config.C.Address); err != nil {
		panic(err.Error())
	}
}
