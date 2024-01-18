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

func GetNotesByUserID(c *fiber.Ctx) error {

	userID, err := helpers.IsValidNumber(c.Params("userid"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("User ID is invalid")
	}

	var notes []domain.Note
	if err := database.DB.Find(&notes, userID).Error; err != nil {
		log.Printf("error while fetching the data %v", err)
		return c.Status(fiber.StatusNoContent).SendString("Phrases are not availble")
	}

	vmNotes := make([]viewmodel.VM_note, len(notes))
	for i, v := range vmNotes {
		vmNotes[i] = viewmodel.VM_note{
			NoteID:   v.NoteID,
			Title:    v.Title,
			Detail:   v.Detail,
			NoteType: v.NoteType,
		}
	}

	result, err := json.Marshal(vmNotes)
	if err != nil {
		log.Printf("JSON marshaling issue, %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	return c.Status(fiber.StatusOK).SendString(string(result))
}

func GetNoteByID(c *fiber.Ctx) error {
	// Validate the input parameter
	noteID, err := helpers.IsValidNumber(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input parameter")
	}

	// Fetch note from the database
	var note domain.Note
	if err := database.DB.Find(&note, noteID).Error; err != nil {
		log.Printf("Phrase not found: %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Converting domain model to view model
	vocab := viewmodel.VM_note{
		NoteID:   note.NoteID,
		Title:    note.Title,
		Detail:   note.Detail,
		NoteType: note.NoteType,
	}

	// Marshal the view model from JSON
	result, err := json.Marshal(vocab)
	if err != nil {
		log.Fatalf("JSON marshaling error %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	return c.Status(fiber.StatusOK).SendString(string(result))
}

func AddNote(c *fiber.Ctx) error {
	note := new(viewmodel.VM_note)
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	vmNote := domain.Note{
		NoteID:   note.NoteID,
		Title:    note.Title,
		Detail:   note.Detail,
		NoteType: note.NoteType,
	}
	if database.DB.Create(&vmNote).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Record added successfully")
}

func EditNote(c *fiber.Ctx) error {
	noteID, err := helpers.IsValidNumber(c.Params("id"))
	if err != nil {
		return c.SendString("")
	}

	note := new(viewmodel.VM_note)
	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	vmPhrase := domain.Note{
		NoteID:   noteID,
		Title:    note.Title,
		Detail:   note.Detail,
		NoteType: note.NoteType,
	}
	if database.DB.Save(&vmPhrase).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("internal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Vocabulary updated successfully")
}

func RemoveNote(c *fiber.Ctx) error {
	noteID, err := helpers.IsValidNumber(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Note id not found")
	}

	if database.DB.Delete(&domain.Note{}, noteID).Error != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Interal server error")
	}
	return c.Status(fiber.StatusOK).SendString("Record deleted successfully")
}
