package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_dashboard(router fiber.Router) {
	router.Get("/", controller.UserLearning)
	router.Get("/", controller.Last10)
	router.Get("/", controller.LearntYesterday)
	router.Get("/", controller.LearntIn7Days)
}
