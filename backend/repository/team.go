package repository

import (
	"backend/types/database"
	"backend/types/extend"
)

type TeamRepository interface {
	GetTeamById(uint64) *database.Team
	GetTeams() []*database.Team
	GetTurned() []*database.Team
	SetTurned([]*database.Team)
	GetLeaderboardConn() *extend.ConnModel
	GetAdminConn() *extend.ConnModel
	GetStudentConns() []*extend.ConnModel
}
