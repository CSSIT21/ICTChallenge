package services

import (
	"backend/loaders/hub"
	"backend/repository"
	"backend/types/message"
	"backend/types/payload"
	"backend/types/response"
)

type topicService struct {
	topicEvent repository.TopicRepository
}

func NewTopicService(topicRepository repository.TopicRepository) topicService {
	return topicService{topicEvent: topicRepository}
}

func (s topicService) OpenCard(body *payload.OpenCard) error {
	topics, err := s.topicEvent.GetTopics()
	if err != nil {
		return &response.Error{
			Message: "Unable to get topics",
		}
	}

	if s.topicEvent.GetCurrentCard() != nil {
		return &response.Error{
			Message: "The card has already opened",
		}
	}

	if topics[body.TopicId-1].Cards[body.CardId-1].Opened {
		return &response.Error{
			Message: "The card has already opened",
		}
	} else {
		topics[body.TopicId-1].Cards[body.CardId-1].Opened = true
	}

	hub.Hub.CardProjectorConn.Emit(&message.OutboundMessage{
		Event: message.CardOpen,
		Payload: map[string]any{
			"card_id":  topics[body.TopicId-1].Cards[body.CardId-1].Id,
			"topic_id": topics[body.TopicId-1].Id,
			"question": topics[body.TopicId-1].Cards[body.CardId-1].Question,
			"bonus":    topics[body.TopicId-1].Cards[body.CardId-1].Bonus,
		},
	})

	s.topicEvent.SetCurrentCard(topics[body.TopicId-1].Cards[body.CardId-1])

	return nil
}
