package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_revise(router fiber.Router) {
	router.Get("/", controller.VocabularyByPeriod)
	router.Get("/", controller.PhraseByPeriod)
	router.Post("/", controller.AddRevision)
	router.Post("/", controller.RevisionHistoryByUserID)
}
