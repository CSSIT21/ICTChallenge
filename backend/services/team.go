package services

import (
	"backend/types"
	"backend/types/database"
)

type TeamService interface {
	GetAllTeams() ([]*database.Team, error)
	GetTeamById(uint64) (*database.Team, error)
	ChangeTeamScore(team *database.Team, problem *database.Card, answer bool, correct bool, bonus bool) error
	GetPodium() ([]*database.Podium, error)
	GetRanking() ([]*database.Ranking, error)
	GetTeamInfo(team *database.Team) (*types.TeamInfo, error)
}
