package router

import (
	"github.com/gofiber/fiber/v2"

	"backend/loaders/hub"
	"backend/repository"
	"backend/services"
)

func Init(router fiber.Router) {
	// * Registrations

	// * Team
	teamRepository := repository.NewTeamEvent(hub.Hub)
	teamService := services.NewTeamService(teamRepository)
	_ = teamService
	// teamHandler := handlers.(rankingService)

	// * Paths

	// * Team
	// ranking := router.Group("team/")
	// ranking.Get("all", teamService)
}
