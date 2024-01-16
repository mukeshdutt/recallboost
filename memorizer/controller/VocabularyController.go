package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

// Add new vocabulary
func AddVocabulary(c *fiber.Ctx) error {

	vocab := new(viewmodel.VM_vocabulary)
	if err := c.BodyParser(vocab); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	vocabulary := domain.Vocabulary{
		VocabularyID:  vocab.VocabularyID,
		Vocabulary:    vocab.Vocabulary,
		Detail:        vocab.Detail,
		ReferenceFrom: vocab.ReferenceFrom,
		UsageType:     vocab.UsageType,
		Antonyms:      vocab.Antonyms,
		Synomyms:      vocab.Synomyms,
		PartsofSpeech: vocab.PartsofSpeech,
	}
	if database.DB.Create(&vocabulary).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Record added successfully")
}

// Edit existing vacabulary
func EditVocabulary(c *fiber.Ctx) error {

	vocabID, err := helpers.IsValidNumber(c.Params("id"))
	if err != nil {
		return c.SendString("")
	}

	vocab := new(viewmodel.VM_vocabulary)
	if err := c.BodyParser(vocab); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	vocabulary := domain.Vocabulary{
		VocabularyID:  vocabID,
		Vocabulary:    vocab.Vocabulary,
		Detail:        vocab.Detail,
		ReferenceFrom: vocab.ReferenceFrom,
		UsageType:     vocab.UsageType,
		Antonyms:      vocab.Antonyms,
		Synomyms:      vocab.Synomyms,
		PartsofSpeech: vocab.PartsofSpeech,
		ModifiedBy:    1,
		ModifiedAt:    time.Now(),
	}
	if database.DB.Save(&vocabulary).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Vocabulary updated successfully")
}

// Remove existing vocabulary
func RemoveVocabulary(c *fiber.Ctx) error {

	vocabID, err := helpers.IsValidNumber(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("vocabulary id not found")
	}

	if database.DB.Delete(&domain.Vocabulary{}, vocabID).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Interal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Record deleted successfully")
}
