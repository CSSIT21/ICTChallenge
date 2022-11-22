package services

import (
	"backend/repository"
	"backend/types/database"
	"backend/types/extend"
	"backend/types/payload"
	"backend/types/response"
)

type topicService struct {
	topicEvent repository.TopicRepository
}

func NewTopicService(topicRepository repository.TopicRepository) *topicService {
	return &topicService{topicEvent: topicRepository}
}

func (s *topicService) OpenCard(body *payload.OpenCard) ([]*database.Topic, error) {
	topics := s.topicEvent.GetTopics()

	if s.topicEvent.GetCurrentCard() != nil {
		return nil, &response.Error{
			Message: "Opened card remaining",
		}
	}

	if topics[body.TopicId-1].Cards[body.CardId-1].Opened {
		return nil, &response.Error{
			Message: "The card has already opened",
		}
	}

	// hub.Hub.CardProjectorConn.Emit(&message.OutboundMessage{
	//	Event: message.CardState,
	//	Payload: map[string]any{
	//		"mode":   "topic",
	//		"topics": updatedTopics,
	//	},
	// })
	//
	// hub.Hub.CardProjectorConn.Emit(&message.OutboundMessage{
	//	Event: message.CardOpen,
	//	Payload: map[string]any{
	//		"card_id":  topics[body.TopicId-1].Cards[body.CardId-1].Id,
	//		"topic_id": topics[body.TopicId-1].Id,
	//		"question": topics[body.TopicId-1].Cards[body.CardId-1].Question,
	//		"bonus":    topics[body.TopicId-1].Cards[body.CardId-1].Bonus,
	//	},
	// })

	s.topicEvent.SetCurrentCard(topics[body.TopicId-1].Cards[body.CardId-1])

	return topics, nil
}

func (s *topicService) GetCardConn() *extend.ConnModel {
	return s.topicEvent.GetCardConn()
}

func (s *topicService) GetTopics() []*database.Topic {
	return s.topicEvent.GetTopics()
}
