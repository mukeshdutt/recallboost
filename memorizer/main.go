package main

import (
	_ "database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/memorizer/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func GetData() {

	dsn := "root:@Pass810@tcp(127.0.0.1:3306)/memorize?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error occurred")
	}
}
