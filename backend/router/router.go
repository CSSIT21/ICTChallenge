package router

import (
	"backend/handlers"
	"backend/repository"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router) {
	// * Register -----------

	// * Ranking
	rankingRepository := repository.NewRankingEvent()
	rankingService := services.NewRankingService(rankingRepository)
	rankingHandler := handlers.NewRankingHandler(rankingService)

	// * Paths --------------
	ranking := router.Group("ranking/")
	ranking.Get("all", rankingHandler.GetAllRankings)
}
