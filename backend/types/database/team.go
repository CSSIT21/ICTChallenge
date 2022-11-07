package database

import "backend/loaders/hub"

type Team struct {
	Id          uint64           `json:"id"`
	Name        string           `json:"name"`
	School      string           `json:"school"`
	Score       int              `json:"score"`
	Token       string           `json:"token"`
	Connections []*hub.ConnModel `json:"-"`
}
