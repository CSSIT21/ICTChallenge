package router

import (
	"backend/handler"
	"github.com/gofiber/fiber/v2"

	"backend/loaders/fiber/middleware"
	"backend/loaders/hub"
	"backend/repository"
	"backend/services"
	"backend/utils/config"
)

func Init(router fiber.Router) {
	// * Registrations

	// * Team
	teamRepository := repository.NewTeamEvent(hub.Hub)
	teamService := services.NewTeamService(teamRepository)
	teamHandler := handler.NewTeamHandler(teamService)

	// * Topic
	topicRepository := repository.NewTopicEvent(hub.Hub)
	topicService := services.NewTopicService(topicRepository)
	_ = topicService

	// * Paths

	// * Admin
	admin := router.Group("am/", middleware.Auth(config.C.AdminSecret))
	admin.Get("info", teamHandler.GetAllTeams)
}
