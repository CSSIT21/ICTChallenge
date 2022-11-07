package repository

import (
	"backend/hub"
	"backend/types/database"
)

type teamEvent struct {
	hub *hub.Model
}

func NewTeamEvent(hub *hub.Model) teamEvent {
	return teamEvent{hub: hub}
}

func (r teamEvent) GetTeamById(id uint64) (*database.Team, error) {
	return nil, nil
}

func (r teamEvent) GetTeams() ([]*database.Team, error) {
	return nil, nil
}

func (r teamEvent) GetPodium() ([]*database.Podium, error) {
	return nil, nil
}

func (r teamEvent) ChangeTeamScore(id uint64, score float64) error {
	return nil
}

func (r teamEvent) GetRanking() ([]*database.Ranking, error) {
	return nil, nil
}
