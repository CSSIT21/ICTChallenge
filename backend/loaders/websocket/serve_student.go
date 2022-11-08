package websocket

import (
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/loaders/hub"
	"backend/types/database"
	"backend/types/extend"
	"backend/types/message"
	"backend/utils/logger"
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
	model := &extend.ConnModel{
		Context: "",
		Conn:    nil,
		Mutex:   nil,
	}

	team.Connections = append(team.Connections, model)

	for {
		t, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if t != websocket.TextMessage {
			break
		}

		model.Emit(&message.OutboundMessage{
			Event:   message.EchoEvent,
			Payload: p,
		})
	}

	// * Close connection
	if err := conn.Close(); err != nil {
		logger.Log(logrus.Warn, "UNHANDLED CONNECTION CLOSE: "+err.Error())
	}

	// * Reset team connection
	var connections []*extend.ConnModel
	for _, connection := range team.Connections {
		if connection != model {
			connections = append(connections, connection)
		}
	}
	team.Connections = connections
}
