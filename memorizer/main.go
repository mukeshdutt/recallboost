package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/routes"
)

func main() {

	app := fiber.New()

	routes.R_dashboard(app.Group("/jfkdfs"))

}
