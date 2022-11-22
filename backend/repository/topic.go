package repository

import (
	"backend/types/database"
	"backend/types/enum"
	"backend/types/extend"
)

type TopicRepository interface {
	GetQuestion(card *database.Card) (*database.Question, error)
	GetTopics() []*database.Topic
	GetCurrentCard() *database.Card
	SetCurrentCard(card *database.Card)
	GetCardConn() *extend.ConnModel
	GetMode() enum.Mode
	SetMode(enum.Mode)
	GetPreviewCount() uint8
	SetPreviewCount(uint8)
}
