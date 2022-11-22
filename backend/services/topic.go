package services

import (
	"backend/types/database"
	"backend/types/extend"
	"backend/types/payload"
)

type TopicService interface {
	OpenCard(body *payload.OpenCard) ([]*database.Topic, error)
	GetCardConn() *extend.ConnModel
	GetTopics() []*database.Topic
}
