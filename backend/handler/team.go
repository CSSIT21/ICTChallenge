package handler

import (
	"backend/services"
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
