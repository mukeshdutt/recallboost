package controller

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/database"
	"github.com/memorizer/domain"
	"github.com/memorizer/helpers"
	"github.com/memorizer/viewmodel"
)

// Show all the phrases
func GetAllPhrase(c *fiber.Ctx) error {

	var phrases []domain.Phrase
	if err := database.DB.Find(&phrases).Error; err != nil {
		log.Printf("error while fetching the data %v", err)
		return c.Status(fiber.StatusNoContent).SendString("Phrases are not availble")
	}

	vmPhrases := make([]viewmodel.VM_phrase, len(phrases))
	for i, v := range vmPhrases {
		vmPhrases[i] = viewmodel.VM_phrase{
			PhraseID:      v.PhraseID,
			Phrase:        v.Phrase,
			Detail:        v.Detail,
			ReferenceFrom: v.ReferenceFrom,
			UsageType:     v.UsageType,
			PhraseType:    v.PhraseType,
		}
	}

	result, err := json.Marshal(vmPhrases)
	if err != nil {
		log.Printf("JSON marshaling issue, %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	return c.Status(fiber.StatusOK).SendString(string(result))
}

// Get Phrases By ID
func GetPhraseByID(c *fiber.Ctx) error {

	// Validate the input parameter
	phraseID, err := helpers.IsValidNumber(c.Params("name"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input parameter")
	}

	// Fetch phrase from the database
	var phrase domain.Phrase
	if err := database.DB.Find(&phrase, phraseID).Error; err != nil {
		log.Printf("Phrase not found: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Converting domain model to view model
	vocab := viewmodel.VM_phrase{
		PhraseID:      phrase.PhraseID,
		Phrase:        phrase.Phrase,
		Detail:        phrase.Detail,
		ReferenceFrom: phrase.ReferenceFrom,
		UsageType:     phrase.UsageType,
	}

	// Marshal the view model from JSON
	result, err := json.Marshal(vocab)
	if err != nil {
		log.Fatalf("JSON marshaling error %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.Status(fiber.StatusOK).SendString(string(result))
}

// Introduce new phrase
func AddPhrase(c *fiber.Ctx) error {

	phrase := new(viewmodel.VM_phrase)
	if err := c.BodyParser(phrase); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	vmPhrase := domain.Phrase{
		PhraseID:      phrase.PhraseID,
		Phrase:        phrase.Phrase,
		Detail:        phrase.Detail,
		ReferenceFrom: phrase.ReferenceFrom,
		UsageType:     phrase.UsageType,
	}
	if database.DB.Create(&vmPhrase).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Record added successfully")
}

// Update existing phrase by id
func EditPhrase(c *fiber.Ctx) error {
	phraseID, err := helpers.IsValidNumber(c.Params("id"))
	if err != nil {
		return c.SendString("")
	}

	phrase := new(viewmodel.VM_phrase)
	if err := c.BodyParser(phrase); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	vmPhrase := domain.Phrase{
		PhraseID:      phraseID,
		Phrase:        phrase.Phrase,
		Detail:        phrase.Detail,
		ReferenceFrom: phrase.ReferenceFrom,
		UsageType:     phrase.UsageType,
	}
	if database.DB.Save(&vmPhrase).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Vocabulary updated successfully")
}

// Remove phrase by id
func RemovePhrase(c *fiber.Ctx) error {
	phraseID, err := helpers.IsValidNumber(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("vocabulary id not found")
	}

	if database.DB.Delete(&domain.Phrase{}, phraseID).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Interal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Record deleted successfully")
}
