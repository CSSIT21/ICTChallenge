package websocket

import (
	"github.com/gofiber/websocket/v2"
	"github.com/sirupsen/logrus"

	"backend/repository"
	"backend/services"

	"backend/loaders/hub"
	"backend/types/extend"
	"backend/types/message"
	"backend/utils/logger"
	"backend/utils/value"
)

func ServeProjector(model *extend.ConnModel, conn *websocket.Conn, token string) {
	// * Validate admin token
	if conn.Query("token") != token {
		return
	}

	hub.HandleConnectionSwitch(model)

	// * Assign connection
	model.Mutex.Lock()
	model.Conn = conn
	model.Mutex.Unlock()

	// * Topic
	topicRepository := repository.NewTopicEvent(hub.Hub)

	// * Team
	teamRepository := repository.NewTeamEvent(hub.Hub)
	teamService := services.NewTeamService(teamRepository, topicRepository)

	// * Initial emit
	if model.Context == "LEADERBOARD_PROJECTOR_CONN" {
		rankings, _ := teamService.GetRanking()
		hub.Hub.LeaderboardProjectorConn.Emit(&message.OutboundMessage{
			Event: message.LeaderboardState,
			Payload: map[string]any{
				"rankings": rankings,
			},
		})
	}

	if model.Context == "CARD_PROJECTOR_CONN" {
		// TODO: Card initial emit
	}

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

	// * Reset player connection
	model.Conn = nil

	// * Unlock in case of connection switch
	if value.MutexLocked(model.Mutex) {
		model.Mutex.Unlock()
	}
}
