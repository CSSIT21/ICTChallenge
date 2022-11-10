package services

import (
	"backend/types/database"
	"backend/types/extend"
	"backend/types/payload"
)

type TeamService interface {
	GetAllTeamInfos() ([]*payload.TeamInfo, error)
	GetCurrentScore(team *database.Team) int32
	GetPodium() ([]*payload.Podium, error)
	GetRanking() ([]*payload.TeamScore, error)
	UpdateScore(body *payload.UpdateScore) ([]*payload.TeamScore, error)
	GetNextTurn() *database.Team
	GetStudentsTurn() *payload.StudentTurn
	GetStudentConns() []*extend.ConnModel
	GetAdminConn() *extend.ConnModel
	GetLeaderBoardConn() *extend.ConnModel
}
