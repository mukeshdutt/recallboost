package controller

import "github.com/gofiber/fiber/v2"

func NoteType(c *fiber.Ctx) error {
	// sampleResponse := "Hello from note controller file"
	// return c.Status(200).JSON(fiber.Map{
	// 	"success": true,
	// })
	return c.SendString("hello response from first notes api")
}

func Notes(c *fiber.Ctx) {

}
