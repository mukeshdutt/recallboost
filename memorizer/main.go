package main

import (
	_ "database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/database"
	"github.com/memorizer/domain"
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

	// ---- database logic checking here
	db := database.Connection()

	var vocabulary domain.Vocabulary
	db.Find(&vocabulary)
	fmt.Println(vocabulary)

	// app.Listen(":3000")
}
