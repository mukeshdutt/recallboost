package main

import (
	_ "database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/routes"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	app := fiber.New()

	routes.R_dashboard(app.Group("/api/dashboard"))
	routes.R_note(app.Group("/api/note"))
	routes.R_phrase(app.Group("/api/phrase"))
	routes.R_quiz(app.Group("/api/quiz"))
	routes.R_revise(app.Group("/api/revise"))
	routes.R_vocabulary(app.Group("/api/vocabulary"))

	app.Listen(":3000")
}
