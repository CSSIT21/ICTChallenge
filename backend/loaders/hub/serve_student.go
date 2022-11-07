package hub

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

func (r *Model) Serve(conn *websocket.Conn) {

	// * Check game entrance availability

	for {
		var (
			mt  int
			msg []byte
			err error
		)
		if mt, msg, err = conn.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		if err = conn.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
