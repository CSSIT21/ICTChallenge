package services

import (
	"backend/types/database"
	"backend/types/enum"
	"backend/types/extend"
	"backend/types/payload"
)

type TeamService interface {
	GetAllTeamInfos() ([]*payload.TeamInfo, error)
	GetCurrentScore(team *database.Team) int32
	GetPodium() []*payload.Podium
	GetRanking() []*payload.TeamScore
	UpdateScore(body *payload.UpdateScore) ([]*payload.TeamScore, error)
	GetNextTurn() *database.Team
	GetStudentsTurn(team *database.Team) *payload.StudentTurn
	GetStudentConns() []*extend.ConnModel
	GetAdminConn() *extend.ConnModel
	GetLeaderboardConn() *extend.ConnModel
	SetMode(mode enum.Mode)
	IncreasePreview()
	GetLastTurn() *database.Team
}
