package services

import (
	"backend/repository"
	"backend/types/database"
)

type topicService struct {
	topicEvent repository.TopicRepository
}

func NewTopicService(topicRepository repository.TopicRepository) topicService {
	return topicService{topicEvent: topicRepository}
}

func (s topicService) OpenCard(card *database.Card) (string, error) {
	a, _ := s.topicEvent.OpenCard(card)
	return a, nil
}
