package controller

import "github.com/gofiber/fiber/v2"

// Show the count of Phrases and Vocabulary learned so far
func UserLearning(c *fiber.Ctx) error {
	return c.SendString("")
}

func Last10(c *fiber.Ctx) error {
	return c.SendString("")
}

func LearntYesterday(c *fiber.Ctx) error {
	return c.SendString("")
}

func LearntIn7Days(c *fiber.Ctx) error {
	return c.SendString("")
}
