package repository

import (
	"backend/loaders/hub"
	"backend/types/database"
	"backend/types/enum"
	"backend/types/extend"
)

type topicEvent struct {
	hub *hub.Model
}

func NewTopicEvent(hub *hub.Model) *topicEvent {
	return &topicEvent{hub: hub}
}

func (r *topicEvent) GetQuestion(card *database.Card) (*database.Question, error) {
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

func (r *topicEvent) SetMode(mode enum.Mode) {
	r.hub.Mode = mode
}

func (r *topicEvent) GetMode() enum.Mode {
	return r.hub.Mode
}

func (r *topicEvent) SetPreviewCount(count uint8) {
	r.hub.PreviewCount = count
}

func (r *topicEvent) GetPreviewCount() uint8 {
	return r.hub.PreviewCount
}

func (r *topicEvent) GetAdminConn() *extend.ConnModel {
	return r.hub.AdminConn
}
