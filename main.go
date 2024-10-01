package main

import (
	"lazycoder-goftn/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./web/views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Static("/public", "./web/public")
	internal.SetupRoutes(app)
	app.Listen(":3000")
}
