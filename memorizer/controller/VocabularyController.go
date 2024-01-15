package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/database"
	"github.com/memorizer/domain"
)

func GetVocabulary(c *fiber.Ctx) error {

	db := database.Connection()

	var vocab = []domain.Vocabulary{}
	db.Find(&vocab)
	fmt.Println(len(vocab))

	return c.SendString("")
}

func GetAllVocabulary(c *fiber.Ctx) error {
	return c.SendString("")
}

func AddVocabulary(c *fiber.Ctx) error {
	return c.SendString("")
}

func EditVocabulary(c *fiber.Ctx) error {
	return c.SendString("")
}

func RemoveVocabulary(c *fiber.Ctx) error {
	return c.SendString("")
}
