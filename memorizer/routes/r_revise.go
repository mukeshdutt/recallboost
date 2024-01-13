package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_revise(route *fiber.Group) {
	route.Get("/", controller.VocabularyByPeriod)
	route.Get("/", controller.PhraseByPeriod)
	route.Post("/", controller.AddRevision)
	route.Post("/", controller.RevisionHistoryByUserID)
}
