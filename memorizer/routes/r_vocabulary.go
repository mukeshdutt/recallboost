package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_vocabulary(router fiber.Router) {
	router.Get("/all", controller.GetAllVocabulary)
	router.Get("/", controller.GetVocabulary)
	router.Post("/", controller.AddVocabulary)
	router.Put("/", controller.EditVocabulary)
	router.Delete("/", controller.RemoveVocabulary)
}
