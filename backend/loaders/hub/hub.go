package hub

import (
	"backend/types/database"
	"backend/types/extend"
)

var Hub *Model

type Model struct {
	Topics                   []*database.Topic `json:"topics"`
	Teams                    []*database.Team  `json:"teams"`
	Turned                   []*database.Team  `json:"-"`
	CurrentCard              *database.Card    `json:"-"`
	AdminConn                *extend.ConnModel `json:"-"`
	LeaderboardProjectorConn *extend.ConnModel `json:"-"`
	CardProjectorConn        *extend.ConnModel `json:"-"`
}
