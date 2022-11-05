package services

import "backend/types/database"

type TeamService interface {
	GetAllTeams() ([]*database.Team, error)
	GetTeamById(uint64) (*database.Team, error)
	ChangeTeamScore(team *database.Team, problem *database.Question, correct bool, bonus bool) error
	GetPodium() ([]*database.Podium, error)
	GetRanking() ([]*database.Ranking, error)
}
