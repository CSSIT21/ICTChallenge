package repository

import "backend/types/database"

type TopicRepository interface {
	OpenCard(card *database.Question) (string, error)
	GetTopics() ([]*database.Topic, error)
}
