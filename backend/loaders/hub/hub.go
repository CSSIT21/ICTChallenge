package hub

import (
	"backend/types/database"
	"backend/types/extend"
)

var Hub *Model

type Model struct {
	Teams                    []*database.Team  `json:"teams"`
	Topics                   []*database.Topic `json:"topics"`
	AdminConn                *extend.ConnModel `json:"-"`
	LeaderboardProjectorConn *extend.ConnModel `json:"-"`
	CardProjectorConn        *extend.ConnModel `json:"-"`
}
