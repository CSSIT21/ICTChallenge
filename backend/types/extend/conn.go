package extend

import (
	"fmt"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/types/message"
	"backend/utils/logger"
)

type ConnModel struct {
	Context string
	Conn    *websocket.Conn
	Mutex   *sync.Mutex
}

func (r *ConnModel) Emit(payload *message.OutboundMessage) {
	if r.Conn == nil || r.Conn.Conn == nil {
		return
	}

	r.Mutex.Lock()
	if err := r.Conn.WriteJSON(payload); err != nil {
		logger.Log(logrus.Warn, fmt.Sprintf("WRITING MESSAGE FAILURE FOR PLAYER %s: %s", r.Context, err.Error()))
	}
	r.Mutex.Unlock()
}
