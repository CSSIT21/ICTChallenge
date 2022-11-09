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

	// * Topic
	topicRepository := repository.NewTopicEvent(hub.Hub)
	topicService := services.NewTopicService(topicRepository)
	topicHandler := handler.NewtopicHandler(topicService)

	// * Team
	teamRepository := repository.NewTeamEvent(hub.Hub)
	teamService := services.NewTeamService(teamRepository, topicRepository)
	teamHandler := handler.NewTeamHandler(teamService)

	// * Paths

	// * Admin
	admin := router.Group("am/", middleware.Auth(config.C.AdminSecret))
	admin.Get("info", teamHandler.GetAllTeams)
	admin.Patch("score", teamHandler.UpdateScore)
	admin.Patch("end", teamHandler.EndGame)

	// * Card
	card := router.Group("st/")
	card.Put("open", topicHandler.OpenCard)
}
