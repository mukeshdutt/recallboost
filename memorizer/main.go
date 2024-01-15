package main

import (
	_ "database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/memorizer/database"
	"github.com/memorizer/routes"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	// Connect to the database
	_, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Setting up fiber
	app := fiber.New()

	// Application routes and controllers
	routes.R_dashboard(app.Group("/api/dashboard"))
	routes.R_note(app.Group("/api/note"))
	routes.R_phrase(app.Group("/api/phrase"))
	routes.R_quiz(app.Group("/api/quiz"))
	routes.R_revise(app.Group("/api/revise"))
	routes.R_vocabulary(app.Group("/api/vocabulary"))

	// Using the middleware for request logging
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, https://gofiber.net",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Starting the fiber application at port 3000
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Error starting fiber", err)
	}
}
