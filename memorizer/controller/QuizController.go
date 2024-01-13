package controller

import "github.com/gofiber/fiber/v2"

func QuizBegin(c *fiber.Ctx) error {
	return c.SendString("")
}

func QuizResult(c *fiber.Ctx) error {
	return c.SendString("")
}

func QuizHistoryByUserID(c *fiber.Ctx) error {
	return c.SendString("")
}
