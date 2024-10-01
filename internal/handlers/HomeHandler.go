package handlers

import "github.com/gofiber/fiber/v2"

func HomeHandler(c *fiber.Ctx) error {
	pageName := "pages/index/index"
	return c.Render(pageName, fiber.Map{
		"Title":  "Home",
		"JsFile": pageName,
	})
}
