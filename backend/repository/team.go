package repository

import (
	"backend/types/database"
)

type TeamRepository interface {
	GetTeamById(uint64) (*database.Team, error)
	GetTeams() ([]*database.Team, error)
	GetTurned() ([]*database.Team, error)
}
