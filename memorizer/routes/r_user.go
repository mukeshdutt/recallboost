package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/controller"
)

func R_user(router fiber.Router) {
	router.Get("", controller.Register)
	router.Get("", controller.Login)
	router.Get("", controller.ForgotPassword)
	router.Get("", controller.ResetPassword)
}
