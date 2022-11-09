package handlers

import (
	"backend/services"
	"backend/types/response"
	"github.com/gofiber/fiber/v2"
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
	return c.JSON(response.New(accessories))
}
