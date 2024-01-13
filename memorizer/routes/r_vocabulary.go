package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_vocabulary(route *fiber.Group) {
	route.Get("/", controller.GetAllVocabulary)
	route.Get("/", controller.GetVocabulary)
	route.Post("/", controller.AddVocabulary)
	route.Put("/", controller.EditVocabulary)
	route.Delete("/", controller.RemoveVocabulary)
}
