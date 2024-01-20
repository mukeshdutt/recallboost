package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_user(router fiber.Router) {
	router.Get("/register", controller.Register)
	router.Get("/login", controller.Login)
	router.Get("/forgot", controller.ForgotPassword)
	router.Get("/reset", controller.ResetPassword)
}
