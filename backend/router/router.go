package router

import (
	"github.com/gofiber/fiber/v2"

	"backend/handler"

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
	admin.Get("info", teamHandler.GetAllTeamInfos)
	admin.Patch("score", teamHandler.UpdateScore)
	admin.Patch("end", teamHandler.EndGame)

	// * Student
	student := router.Group("st/", middleware.Auth(config.C.StudentSecret))
	student.Put("open", topicHandler.OpenCard)
	student.Get("info", teamHandler.GetTeam)

	// * Card
	// card := router.Group("cd/")
	// card.Get("state")
	// card.Get("open")
}
