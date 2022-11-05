package router

import (
	"backend/hub"
	"github.com/gofiber/fiber/v2"

	"backend/handlers"
	"backend/repository"
	"backend/services"
)

func Init(router fiber.Router) {
	// * Registrations
	
	// Ranking
	rankingRepository := repository.NewRankingEvent(hub.Hub)
	rankingService := services.NewRankingService(rankingRepository)
	rankingHandler := handlers.NewRankingHandler(rankingService)

	// * Paths

	// Ranking
	ranking := router.Group("ranking/")
	ranking.Get("all", rankingHandler.GetAllRankings)
}
