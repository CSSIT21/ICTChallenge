package websocket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"backend/loaders/hub"
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

	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	router.Get("/student", websocket.New(func(conn *websocket.Conn) {
		hub.Hub.Serve(conn)
	}))

	router.Get("/projector", websocket.New(func(conn *websocket.Conn) {

	}))
}
