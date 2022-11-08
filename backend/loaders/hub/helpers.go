package hub

import (
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/types/extend"
	"backend/types/message"
	"backend/utils/logger"
	"backend/utils/value"
)

func HandleConnectionSwitch(conn *extend.ConnModel) {
	// * Connection switch
	if conn.Conn != nil {
		conn.Emit(&message.OutboundMessage{
			Event:   message.ConnectionSwitchEvent,
			Payload: nil,
		})

		conn.Mutex.Lock()
		if err := value.ErrorChain(
			conn.Conn.WriteMessage(websocket.CloseMessage, []byte{}),
			conn.Conn.Close(),
		); err != nil {
			logger.Log(logrus.Warn, "UNHANDLED PLAYER CONN SWITCH: "+err.Error())
		}
	}
}
