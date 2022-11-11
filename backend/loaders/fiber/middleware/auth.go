package middleware

import (
	"github.com/gofiber/fiber/v2"

	"backend/types/response"
)

var Auth = func(token string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// * Check authorization header
		if ctx.Get("Authorization") != "Bearer "+token {
			return ctx.Status(fiber.StatusUnauthorized).JSON(response.New("Unauthorized"))
		}

		return ctx.Next()
	}
}
