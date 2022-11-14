package handler

import (
	"github.com/gofiber/fiber/v2"

	"backend/loaders/hub"
	"backend/services"
	"backend/types/message"
	"backend/types/payload"
	"backend/types/response"
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

	next := h.teamService.GetNextTurn()

	h.teamService.GetLeaderboardConn().Emit(&message.OutboundMessage{
		Event: message.LeaderboardState,
		Payload: map[string]any{
			"highlighted_id": next.Id,
			"rankings":       rankings,
		},
	})

	// Next turn
	student := h.teamService.GetStudentsTurn(next)
	for _, conn := range h.teamService.GetStudentConns() {
		conn.Emit(&message.OutboundMessage{
			Event:   message.StudentTurn,
			Payload: student,
		})
	}

	return c.JSON(response.New("Successfully updated score record"))
}

func (h *teamHandler) EndGame(c *fiber.Ctx) error {
	rankings := h.teamService.GetPodium()

	h.teamService.GetLeaderboardConn().Emit(&message.OutboundMessage{
		Event: message.LeaderboardPodium,
		Payload: map[string]any{
			"rankings": rankings,
		},
	})

	return c.JSON(response.New("Game has ended"))
}

func (h *teamHandler) SetLeaderboardMode(c *fiber.Ctx) error {
	var body *payload.LeaderboardMode
	if err := c.BodyParser(&body); err != nil {
		return &response.Error{
			Message: "Unable to parse body",
			Err:     err,
		}
	}

	h.teamService.SetMode(*body.Mode)

	return c.JSON(response.New("Successfully updated leaderboard mode"))
}

func (h *teamHandler) IncrementPreview(c *fiber.Ctx) error {

	return c.JSON(response.New("Successfully updated leaderboard mode"))
}

func (h *teamHandler) DismissCard(c *fiber.Ctx) error {
	hub.Skip <- true
	return c.JSON(response.New("Successfully dismiss the card"))
}

func (h *teamHandler) PauseCard(c *fiber.Ctx) error {
	hub.Pause <- true
	return c.JSON(response.New("Successfully toggle card pausing"))
}
