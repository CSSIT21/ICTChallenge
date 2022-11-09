package repository

import "backend/types/database"

type TopicRepository interface {
	OpenCard(card *database.Card) (string, error)
	GetTopics() ([]*database.Topic, error)
	GetCurrentCard() *database.Card
	SetCurrentCard(card *database.Card)
}
