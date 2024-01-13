package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func NoteRoutes(route *fiber.App) {
	api := route.Group("/notes")
	api.Get("/notes", controller.AddNote)
}
