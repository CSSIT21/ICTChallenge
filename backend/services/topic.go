package services

import "backend/types/database"

type TopicService interface {
	OpenCard(card *database.Card) (string, error)
}
