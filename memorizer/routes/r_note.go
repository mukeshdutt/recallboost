package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_note(router fiber.Router) {
	router.Get("/", controller.GetNoteByID)
	router.Get("/", controller.GetNotesByUserID)
	router.Post("/", controller.AddNote)
	router.Put("/", controller.EditNote)
	router.Delete("/", controller.RemoveNote)
}
