package repository

import (
	"math/rand"

	"backend/loaders/hub"
	"backend/types/database"
	"backend/utils/value"
)

type teamEvent struct {
	hub *hub.Model
}

func NewTeamEvent(hub *hub.Model) *teamEvent {
	return &teamEvent{hub: hub}
}

func (r *teamEvent) GetTeamById(id uint64) (*database.Team, error) {
	return nil, nil
}

func (r *teamEvent) GetTeams() ([]*database.Team, error) {
	return hub.Hub.Teams, nil
}

func (r *teamEvent) GetTurned() ([]*database.Team, error) {
	return hub.Hub.Turned, nil
}

func (r *teamEvent) GetNextTurn() (*database.Team, error) {
	if len(r.hub.Turned) == 6 {
		r.hub.Turned = r.hub.Turned[:0]
	}

	var candidates []*database.Team
	for _, team := range r.hub.Teams {
		if !value.Contain(r.hub.Turned, team) {
			candidates = append(candidates, team)
		}
	}

	selected := candidates[rand.Intn(len(candidates))]
	r.hub.Turned = append(r.hub.Turned, selected)
	
	return selected, nil
}
