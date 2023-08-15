package main

import (
	"log"

	"github.com/Garry028/trivia/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
