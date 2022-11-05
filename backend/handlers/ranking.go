package handlers

import (
	"github.com/gofiber/fiber/v2"

	"backend/hub"
	"backend/services"
	"backend/types/response"
)

type rankingHandler struct {
	rankingService services.RankingService
}

func NewRankingHandler(rankingService services.RankingService) rankingHandler {
	return rankingHandler{rankingService: rankingService}
}

func (h rankingHandler) GetAllRankings(c *fiber.Ctx) error {
	// token := (*jwt.Token) here
	accessories, err := h.rankingService.GetAllRankings()
	if err != nil {
		return err
	}

	hub.Conn.WriteJSON(map[string]any{
		"hello": "world",
	})

	return c.JSON(response.New(accessories))
}
