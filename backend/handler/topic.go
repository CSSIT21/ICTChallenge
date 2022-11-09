package handler

import (
	"backend/services"
	"backend/types/payload"
	"backend/types/response"
	"github.com/gofiber/fiber/v2"
)

type topicHandler struct {
	teamService services.TopicService
}

func NewtopicHandler(topicService services.TopicService) topicHandler {
	return topicHandler{teamService: topicService}
}

func (h topicHandler) OpenCard(c *fiber.Ctx) error {
	var body *payload.OpenCard
	if err := c.BodyParser(&body); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	if err := h.teamService.OpenCard(body); err != nil {
		return &response.Error{
			Err: err,
		}
	}

	return c.JSON(response.New("The card has opened"))
}
