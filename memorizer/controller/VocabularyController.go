package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/database"
	"github.com/memorizer/domain"
	"github.com/memorizer/viewmodel"
)

func GetVocabulary(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var vocabs []domain.Vocabulary
	database.DB.Find(&vocabs)
	fmt.Println(vocabs)

	if err := database.DB.Find(&vocabs).Error; err != nil {
		log.Printf("Error fetching vocabulary: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Conver domain model to view model
	vmVocabs := make([]viewmodel.VM_vocabulary, len(vocabs))
	for i, vocab := range vocabs {
		vmVocabs[i] = viewmodel.VM_vocabulary{
			VocabularyID:  vocab.VocabularyID,
			Vocabulary:    vocab.Vocabulary,
			Detail:        vocab.Detail,
			ReferenceFrom: vocab.ReferenceFrom,
			UsageType:     vocab.UsageType,
			Antonyms:      vocab.Antonyms,
			Synomyms:      vocab.Synomyms,
			PartsofSpeech: vocab.PartsofSpeech,
		}
	}

	// Marshal the view model to JSON
	result, err := json.Marshal(vmVocabs)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.Status(fiber.StatusOK).SendString(string(result))
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
