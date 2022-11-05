package websocket

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"backend/hub"
)

func Init(router fiber.Router) {
	router.Use("/", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	router.Get("/student", websocket.New(func(conn *websocket.Conn) {
		// c.Locals are added to the *websocket.Conn
		// log.Println(conn.Locals("allowed"))  // true
		// log.Println(conn.Params("id"))       // 123
		// log.Println(conn.Query("v"))         // 1.0
		// log.Println(conn.Cookies("session")) // ""
		fmt.Println(111)

		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		hub.Hub.Serve(conn)
	}))
}
