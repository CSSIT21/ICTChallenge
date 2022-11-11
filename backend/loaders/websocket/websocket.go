package websocket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"

	"backend/loaders/hub"
	"backend/utils/config"
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

	router.Get("/projector/leaderboard", websocket.New(func(conn *websocket.Conn) {
		ServeProjector(hub.Hub.LeaderboardProjectorConn, conn, config.C.ProjectorSecret)
	}))

	router.Get("/projector/card", websocket.New(func(conn *websocket.Conn) {
		ServeProjector(hub.Hub.CardProjectorConn, conn, config.C.ProjectorSecret)
	}))

	router.Get("/admin", websocket.New(func(conn *websocket.Conn) {
		ServeProjector(hub.Hub.AdminConn, conn, config.C.AdminSecret)
	}))

	router.Get("/student", websocket.New(func(conn *websocket.Conn) {
		ServeStudent(conn)
	}))
}
