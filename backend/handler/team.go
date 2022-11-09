package handler

import (
	"backend/loaders/hub"
	"backend/services"
	"backend/types/message"
	"backend/types/payload"
	"backend/types/response"
	"github.com/gofiber/fiber/v2"
)

type teamHandler struct {
	teamService services.TeamService
}

func NewTeamHandler(teamService services.TeamService) teamHandler {
	return teamHandler{teamService: teamService}
}

func (h teamHandler) GetAllTeams(c *fiber.Ctx) error {
	teams, err := h.teamService.GetAllTeams()
	if err != nil {
		return err
	}
	return c.JSON(response.New(teams))
}

func (h teamHandler) UpdateScore(c *fiber.Ctx) error {
	var body *payload.UpdateScore
	if err := c.BodyParser(&body); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	rankings, err := h.teamService.UpdateScore(body)
	if err != nil {
		return err
	}

	hub.Hub.LeaderboardProjectorConn.Emit(&message.OutboundMessage{
		Event: message.LeaderboardState,
		Payload: map[string]any{
			"rankings": rankings,
		},
	})

	return c.JSON(response.New("Score updated"))
}
