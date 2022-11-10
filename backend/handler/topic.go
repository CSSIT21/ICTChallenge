package handler

import (
	"backend/services"
	"backend/types/message"
	"backend/types/payload"
	"backend/types/response"
	"github.com/gofiber/fiber/v2"
)

type topicHandler struct {
	topicService services.TopicService
}

func NewtopicHandler(topicService services.TopicService) topicHandler {
	return topicHandler{topicService: topicService}
}

func (h topicHandler) OpenCard(c *fiber.Ctx) error {
	var body *payload.OpenCard
	if err := c.BodyParser(&body); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	updatedTopics, topics, err := h.topicService.OpenCard(body)
	if err != nil {
		return &response.Error{
			Err: err,
		}
	}

	// CardState
	h.topicService.GetCardConn().Emit(&message.OutboundMessage{
		Event: message.CardState,
		Payload: map[string]any{
			"mode":   "topic",
			"topics": updatedTopics,
		},
	})

	// CardOpen
	h.topicService.GetCardConn().Emit(&message.OutboundMessage{
		Event: message.CardOpen,
		Payload: map[string]any{
			"card_id":  topics[body.TopicId-1].Cards[body.CardId-1].Id,
			"topic_id": topics[body.TopicId-1].Id,
			"question": topics[body.TopicId-1].Cards[body.CardId-1].Question,
			"bonus":    topics[body.TopicId-1].Cards[body.CardId-1].Bonus,
		},
	})

	return c.JSON(response.New("The card has opened"))
}
