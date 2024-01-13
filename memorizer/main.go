package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber"
	"github.com/memorizer/routes"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello, world")
	})

	// app.Get("/hello", Hello)
	// app.Get("/typelist", controller.)

	routes.NoteRoutes(app)
	app.Listen(":4000")
}
