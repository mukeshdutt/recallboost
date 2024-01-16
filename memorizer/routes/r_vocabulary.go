package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_vocabulary(router fiber.Router) {
	router.Get("/all", controller.GetAllVocabulary)
	router.Get("/:id", controller.GetVocabularyByID)
	router.Post("/", controller.AddVocabulary)
	router.Put("/:id", controller.EditVocabulary)
	router.Delete("/:id", controller.RemoveVocabulary)
}
