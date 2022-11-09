package repository

import (
	"backend/loaders/hub"
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
	return hub.Hub.Teams, nil
}

func (r teamEvent) GetTurned() ([]*database.Team, error) {
	return hub.Hub.Turned, nil
}
