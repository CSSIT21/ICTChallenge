package fiber

import (
	"backend/loaders/fiber/middleware"
	"backend/router"
	"backend/types/response"
	"backend/utils/config"
	"github.com/gofiber/fiber/v2"
	"time"
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

	// Register not found handler
	app.Use(notfoundHandler)

	// Startup
	err := app.Listen(config.C.Address)
	if err != nil {
		panic(err.Error())
	}
}
