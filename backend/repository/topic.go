package repository

import (
	"backend/types/database"
	"backend/types/extend"
)

type TopicRepository interface {
	GetQuestion(card *database.Card) (string, error)
	GetTopics() []*database.Topic
	GetCurrentCard() *database.Card
	SetCurrentCard(card *database.Card)
	GetCardConn() *extend.ConnModel
}
