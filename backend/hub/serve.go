package hub

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

var Conn *websocket.Conn

func (r *Model) Serve(conn *websocket.Conn) {
	// * Check game entrance availability
	Conn = conn

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
		log.Printf("mt: %d", mt)
		log.Printf("recv: %s", msg)

		if err = conn.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}
}
