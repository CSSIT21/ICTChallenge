package handler

import (
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

func (h *teamHandler) GetTeam(c *fiber.Ctx) error {

	return nil
}

func (h *teamHandler) GetAllTeamInfos(c *fiber.Ctx) error {
	teams, err := h.teamService.GetAllTeamInfos()
	if err != nil {
		return err
	}
	return c.JSON(response.New(teams))
}

func (h *teamHandler) UpdateScore(c *fiber.Ctx) error {
	var body *payload.UpdateScore
	if err := c.BodyParser(&body); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	// Update score
	rankings, err := h.teamService.UpdateScore(body)
	if err != nil {
		return err
	}

	h.teamService.GetLeaderBoardConn().Emit(&message.OutboundMessage{
		Event: message.LeaderboardState,
		Payload: map[string]any{
			"rankings": rankings,
		},
	})

	// Next turn
	turn := h.teamService.GetStudentsTurn()
	for _, conn := range h.teamService.GetStudentConns() {
		conn.Emit(&message.OutboundMessage{
			Event:   message.StudentTurn,
			Payload: turn,
		})
	}

	return c.JSON(response.New("Successfully updated score record"))
}

func (h *teamHandler) EndGame(c *fiber.Ctx) error {
	rankings, err := h.teamService.GetPodium()
	if err != nil {
		return err
	}

	h.teamService.GetLeaderBoardConn().Emit(&message.OutboundMessage{
		Event: message.LeaderboardPodium,
		Payload: map[string]any{
			"rankings": rankings,
		},
	})

	return c.JSON(response.New("Game has ended"))
}
