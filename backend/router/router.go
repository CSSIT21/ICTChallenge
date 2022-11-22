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
	topicHandler := handler.NewTopicHandler(topicService)

	// * Team
	teamRepository := repository.NewTeamEvent(hub.Hub)
	teamService := services.NewTeamService(teamRepository, topicRepository)
	teamHandler := handler.NewTeamHandler(teamService, topicService)

	// * Paths

	// * Admin
	admin := router.Group("am/", middleware.Auth(config.C.AdminSecret))
	admin.Get("info", teamHandler.GetAllTeamInfos)
	admin.Patch("preview/increment", teamHandler.IncrementPreview)
	admin.Patch("highlight", teamHandler.Highlight)
	admin.Patch("score", teamHandler.UpdateScore)
	admin.Patch("end", teamHandler.EndGame)
	admin.Patch("mode", teamHandler.SetLeaderboardMode)
	admin.Patch("card/skip", teamHandler.SkipCard)
	admin.Patch("card/pause", teamHandler.PauseCard)
	admin.Patch("card/dismiss", teamHandler.DismissCard)

	// * Student
	student := router.Group("st/", middleware.Auth(config.C.StudentSecret))
	student.Put("open", topicHandler.OpenCard)
	student.Get("info", teamHandler.GetTeam)

	// * Card
	// card := router.Group("cd/")
	// card.Get("state")
	// card.Get("open")
}
