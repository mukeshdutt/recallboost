package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_phrase(router fiber.Router) {
	router.Get("/", controller.GetAllPhrase)
	router.Get("/", controller.GetPhrase)
	router.Post("/", controller.AddPhrase)
	router.Put("/", controller.EditPhrase)
	router.Delete("/", controller.RemovePhrase)
}
