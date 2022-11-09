package services

import (
	"backend/types/database"
	"backend/types/payload"
)

type TeamService interface {
	GetAllTeams() ([]*database.Team, error)
	GetTeamById(uint64) (*database.Team, error)
	GetPodium() ([]*payload.Podium, error)
	GetRanking() ([]*payload.TeamInfo, error)
	UpdateScore(body *payload.UpdateScore) ([]*payload.TeamInfo, error)
}
