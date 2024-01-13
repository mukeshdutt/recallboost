package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_dashboard(route *fiber.Group) {
	route.Get("/", controller.UserLearning)
	route.Get("/", controller.Last10)
	route.Get("/", controller.LearntYesterday)
	route.Get("/", controller.LearntIn7Days)
}
