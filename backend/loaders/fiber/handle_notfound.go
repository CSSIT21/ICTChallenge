package fiber

import (
	"backend/types/response"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func notfoundHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
		Success: false,
		Message: fmt.Sprintf("Not found %s", ctx.Path()),
		Error:   "404_NOT_FOUND",
	})
}
