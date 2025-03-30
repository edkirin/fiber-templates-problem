package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func HandleGetFrontpage(c *fiber.Ctx) error {
	fmt.Println("Rendering frontpage")
	return c.Render("frontpage", fiber.Map{})
}

func HandleGetAbout(c *fiber.Ctx) error {
	fmt.Println("Rendering about")
	return c.Render("about", fiber.Map{})
}

func main() {
	engine := html.New("./views", ".tpl")
	engine.Reload(true) // Enable template reload in development mode

	app := fiber.New(fiber.Config{
		Views:         engine,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Get("/", HandleGetFrontpage)
	app.Get("/about", HandleGetAbout)

	// Start the server
	app.Listen(":8000")
}
