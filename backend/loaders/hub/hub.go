package hub

import (
	"backend/types/database"
)

var Hub *Model

type Model struct {
	Teams                    []*database.Team  `json:"teams"`
	Topics                   []*database.Topic `json:"topics"`
	AdminConn                *ConnModel        `json:"admin_conn"`
	LeaderboardProjectorConn *ConnModel        `json:"leaderboard_projector_conn"`
	CardProjectorConn        *ConnModel        `json:"card_projector_conn"`
}
