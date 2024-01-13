package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_quiz(route *fiber.Group) {
	route.Get("/", controller.QuizBegin)
	route.Get("/", controller.QuizResult)
	route.Get("/", controller.QuizHistoryByUserID)
}
