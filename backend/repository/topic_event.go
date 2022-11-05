package repository

import (
	"backend/hub"
	"backend/types/database"
)

type topicEvent struct {
	hub *hub.Model
}

func NewTopicEvent(hub *hub.Model) topicEvent {
	return topicEvent{hub: hub}
}

func (r topicEvent) OpenCard(card *database.Question) (string, error) {
	return card.Problem, nil
}

func (r topicEvent) GetTopics() ([]*database.Topic, error) {
	return []*database.Topic{}, nil
}
