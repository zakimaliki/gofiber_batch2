package main

import (
	"gofiber-batch2/src/configs"
	"gofiber-batch2/src/helpers"
	"gofiber-batch2/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	configs.InitDB()
	helpers.Migration()
	routes.Router(app)
	app.Listen(":3000")
}
