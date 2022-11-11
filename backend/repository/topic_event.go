package repository

import (
	"backend/loaders/hub"
	"backend/types/database"
	"backend/types/extend"
)

type topicEvent struct {
	hub *hub.Model
}

func NewTopicEvent(hub *hub.Model) *topicEvent {
	return &topicEvent{hub: hub}
}

func (r *topicEvent) GetQuestion(card *database.Card) (string, error) {
	return card.Question, nil
}

func (r *topicEvent) GetTopics() []*database.Topic {
	return r.hub.Topics
}

func (r *topicEvent) GetCurrentCard() *database.Card {
	return r.hub.CurrentCard
}

func (r *topicEvent) SetCurrentCard(card *database.Card) {
	r.hub.CurrentCard = card
}

func (r *topicEvent) GetCardConn() *extend.ConnModel {
	return hub.Hub.CardProjectorConn
}
