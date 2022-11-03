package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"

	"backend/types/response"
)

var limiterReached = &response.Error{
	Message: "Rate limit reached, try again in a few minutes.",
}

var Limiter = func() fiber.Handler {
	config := limiter.Config{
		Next:       nil,
		Max:        6000,
		Expiration: 600 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return limiterReached
		},
	}

	return limiter.New(config)
}()
