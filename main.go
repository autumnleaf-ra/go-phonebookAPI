package main

import (
	"github.com/autumnleaf-ra/phonebook-api/apps"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	apps.InitializeApp(app)
}
