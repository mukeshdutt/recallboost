package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/routes"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world")
	})

	app.Get("/hello", Hello)
	// app.Get("/typelist", controllers.TypeList)

	routes.NoteRoutes(app)
	app.Listen(":4000")
}

func Hello(c *fiber.Ctx) error {
	// c.SendString()
	return c.SendString("hello from main file")
}
