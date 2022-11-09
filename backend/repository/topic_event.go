package repository

import (
	"backend/loaders/hub"
	"backend/types/database"
)

type topicEvent struct {
	hub *hub.Model
}

func NewTopicEvent(hub *hub.Model) topicEvent {
	return topicEvent{hub: hub}
}

func (r topicEvent) OpenCard(card *database.Card) (string, error) {
	return card.Question, nil
}

func (r topicEvent) GetTopics() ([]*database.Topic, error) {
	return r.hub.Topics, nil
}

func (r topicEvent) GetCurrentCard() *database.Card {
	return r.hub.CurrentCard
}

func (r topicEvent) SetCurrentCard(card *database.Card) {
	r.hub.CurrentCard = card
}
