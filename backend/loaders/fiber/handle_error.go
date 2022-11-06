package fiber

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"backend/types/response"
)

func errorHandler(ctx *fiber.Ctx, ierr error) error {
	// Case of *fiber.Error.
	if err, ok := ierr.(*fiber.Error); ok {
		return ctx.Status(err.Code).JSON(response.ErrorResponse{
			Success: false,
			Code:    strings.ReplaceAll(strings.ToUpper(err.Error()), " ", "_"),
			Message: err.Error(),
			Data:    nil,
			Error:   err.Error(),
		})
	}

	if err, ok := ierr.(*response.Error); ok {
		// * Apply code fallback
		if err.Code == "" {
			err.Code = "GENERIC_ERROR"
		}

		// * Apply data merging
		var data any
		if err.Detail != "" {
			data = map[string]string{
				"detail": err.Message,
			}
		}

		// * Apply error merging
		var e string
		if err.Err != nil {
			e = err.Err.Error()
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Success: false,
			Code:    err.Code,
			Message: err.Message,
			Data:    data,
			Error:   e,
		})
	}

	// Case of validator.ValidationErrors
	if e, ok := ierr.(validator.ValidationErrors); ok {
		var lists []string
		for _, err := range e {
			lists = append(lists, err.Field()+" ("+err.Tag()+")")
		}

		message := strings.Join(lists[:], ", ")

		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Success: false,
			Code:    "VALIDATION_FAILED",
			Message: "Validation failed on field " + message,
			Error:   e.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(
		response.ErrorResponse{
			Success: false,
			Code:    "UNKNOWN_SERVER_SIDE_ERROR",
			Message: "Unknown server side error",
			Error:   ierr.Error(),
		},
	)
}
