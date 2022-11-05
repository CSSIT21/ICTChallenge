package database

import "github.com/gofiber/websocket/v2"

type Team struct {
	Id     uint64            `json:"id"`
	Name   string            `json:"name"`
	School string            `json:"school"`
	Score  int               `json:"score"`
	Token  string            `json:"token"`
	Conn   []*websocket.Conn `json:"-"`
}
