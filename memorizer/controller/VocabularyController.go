package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/database"
	"github.com/memorizer/domain"
	"github.com/memorizer/helpers"
	"github.com/memorizer/viewmodel"
)

// GetAllVocabulary is use to get list of vocabularies
func GetAllVocabulary(c *fiber.Ctx) error {

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

// Here you can specific vocabulary by the id
func GetVocabularyByID(c *fiber.Ctx) error {

	// Validate the input parameter
	vocabID, err := helpers.IsValidNumber(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input parameter")
	}

	// Fetch vocabulary from the database
	var vocabulary domain.Vocabulary
	if err := database.DB.Find(&vocabulary, vocabID).Error; err != nil {
		log.Printf("Vocabulary not found: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Converting domain model to view model
	vocab := viewmodel.VM_vocabulary{
		VocabularyID:  vocabulary.VocabularyID,
		Vocabulary:    vocabulary.Vocabulary,
		Detail:        vocabulary.Detail,
		ReferenceFrom: vocabulary.ReferenceFrom,
		UsageType:     vocabulary.UsageType,
		Antonyms:      vocabulary.Antonyms,
		Synomyms:      vocabulary.Synomyms,
		PartsofSpeech: vocabulary.PartsofSpeech,
	}

	// Marshal the view model from JSON
	result, err := json.Marshal(vocab)
	if err != nil {
		log.Fatalf("JSON marshaling error %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.Status(fiber.StatusOK).SendString(string(result))
}

func AddVocabulary(c *fiber.Ctx) error {

	c.Request().Body()

	return c.SendString("")
}

func EditVocabulary(c *fiber.Ctx) error {
	return c.SendString("")
}

func RemoveVocabulary(c *fiber.Ctx) error {
	return c.SendString("")
}
