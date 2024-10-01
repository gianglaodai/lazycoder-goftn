package internal

import (
	"github.com/gofiber/fiber/v2"
	"lazycoder-goftn/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
	app.Static("/assets", "../web/assets")
	app.Get("/", handlers.HomeHandler)
}
