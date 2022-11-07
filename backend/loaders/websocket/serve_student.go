package websocket

import (
	"github.com/gofiber/websocket/v2"

	"backend/loaders/hub"
	"backend/types/database"
)

func ServeStudent(conn *websocket.Conn) {
	for _, team := range hub.Hub.Teams {
		if team.Token == conn.Query("token") {
			ServeTeam(team, conn)
			return
		}
	}
}

func ServeTeam(team *database.Team, conn *websocket.Conn) {

}
