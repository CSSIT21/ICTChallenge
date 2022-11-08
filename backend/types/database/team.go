package database

import (
	"backend/types/extend"
)

type Team struct {
	Id          uint64              `json:"id"`
	Name        string              `json:"name"`
	School      string              `json:"school"`
	Token       string              `json:"token"`
	Scores      []*Score            `json:"scores"`
	Connections []*extend.ConnModel `json:"-"`
}
