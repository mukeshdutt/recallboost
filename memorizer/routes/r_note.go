package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_note(route *fiber.Group) {
	route.Get("/", controller.GetNoteByID)
	route.Get("/", controller.GetNotesByUserID)
	route.Post("/", controller.AddNote)
	route.Put("/", controller.EditNote)
	route.Delete("/", controller.RemoveNote)
}
