package repository

import (
	"backend/hub"
	"backend/types/database"
)

type rankingEvent struct {
	hub *hub.Model
}

func NewRankingEvent(hub *hub.Model) rankingEvent {
	return rankingEvent{hub: hub}
}

func (r rankingEvent) GetAll() ([]database.Ranking, error) {
	return []database.Ranking{
		{
			Id:    1,
			Name:  "Mixko",
			Score: 234,
		},
		{
			Id:    2,
			Name:  "Also Mixko",
			Score: 235,
		},
		{
			Id:    3,
			Name:  "Also Mixko",
			Score: 236,
		},
	}, nil
}
