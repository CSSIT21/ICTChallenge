package services

import "backend/types/database"

type TopicService interface {
	OpenCard(card *database.Question) (string, error)
}
