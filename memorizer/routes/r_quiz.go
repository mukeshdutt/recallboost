package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_quiz(router fiber.Router) {
	router.Get("/", controller.QuizBegin)
	router.Get("/", controller.QuizResult)
	router.Get("/", controller.QuizHistoryByUserID)
}
