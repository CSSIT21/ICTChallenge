package mappers

import (
	"backend/types/database"
	"backend/types/response"
	"backend/utils/value"
)

func DisplayTopic(topics []*database.Topic) []*database.Topic {
	var updatedTopics []*database.Topic
	for _, topic := range topics {
		cards, _ := value.Iterate(topic.Cards, func(card *database.Card) (*database.Card, *response.Error) {
			return &database.Card{
				Id:     card.Id,
				Score:  card.Score,
				Opened: card.Opened,
			}, nil
		})
		updated := &database.Topic{
			Id:    topic.Id,
			Title: topic.Title,
			Cards: cards,
		}
		updatedTopics = append(updatedTopics, updated)
	}
	return updatedTopics
}
