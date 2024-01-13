package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_phrase(route *fiber.Group) {
	route.Get("/", controller.GetAllPhrase)
	route.Get("/", controller.GetPhrase)
	route.Post("/", controller.AddPhrase)
	route.Put("/", controller.EditPhrase)
	route.Delete("/", controller.RemovePhrase)
}
