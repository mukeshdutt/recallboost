package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_revise(router fiber.Router) {
	router.Get("/vocabulary", controller.VocabularyByPeriod)
	router.Get("/phrase", controller.PhraseByPeriod)
	router.Post("/", controller.AddRevision)
	router.Post("/history", controller.RevisionHistoryByUserID)
}
