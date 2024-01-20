package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_quiz(router fiber.Router) {
	router.Get("/question", controller.QuizBegin)
	router.Get("/result/:userid", controller.QuizResult)
	router.Get("/history/:userid", controller.QuizHistoryByUserID)
}
