package websocket

import (
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/loaders/hub"
	"backend/types/extend"
	"backend/types/message"
	"backend/utils/logger"
)

func ServeStudent(conn *websocket.Conn) {
	model := &extend.ConnModel{
		Context: "STUDENT_CONN",
		Conn:    conn,
		Mutex:   &sync.Mutex{},
	}

	hub.Hub.StudentConns = append(hub.Hub.StudentConns, model)

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

	// * Update student connection
	var connections []*extend.ConnModel
	for _, connection := range hub.Hub.StudentConns {
		if connection != model {
			connections = append(connections, connection)
		}
	}
	hub.Hub.StudentConns = connections
}
