package repository

import "backend/types/database"

type TeamRepository interface {
	GetTeamById(uint64) (*database.Team, error)
	ChangeTeamScore(uint64, int32) error
	GetTeams() ([]*database.Team, error)
	GetPodium() ([]*database.Podium, error)
	GetRanking() ([]*database.Ranking, error)
}
