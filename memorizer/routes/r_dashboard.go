package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_dashboard(router fiber.Router) {
	router.Get("/user-learning", controller.UserLearning)
	router.Get("/last-10", controller.Last10)
	router.Get("/learnt-yesterday", controller.LearntYesterday)
	router.Get("/learnt-in-7-day", controller.LearntIn7Days)
}
