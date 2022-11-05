package router

import (
	"backend/hub"
	"backend/repository"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router) {
	// * Registrations

	// * Team
	teamRepository := repository.NewTeamEvent(hub.Hub)
	teamService := services.NewTeamService(teamRepository)
	_ = teamService
	//teamHandler := handlers.(rankingService)

	// * Paths

	// * Team
	//ranking := router.Group("team/")
	//ranking.Get("all", teamService)
}
